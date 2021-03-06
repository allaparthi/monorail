package calendar

import (
	"bytes"
	"fmt"
	"html/template"
	"infra/appengine/rotang"
	"net/http"
	"regexp"
	"strings"
	"time"

	"context"

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/server/router"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/api/calendar/v3"
	gcal "google.golang.org/api/calendar/v3"
)

// Calendar implements the rotang.Calenderer interface.
type Calendar struct {
	credentials func(*router.Context) (*http.Client, error)
}

var (
	_      rotang.Calenderer = &Calendar{}
	pregex                   = regexp.MustCompile(`^([.[:alnum:]-_]+) is primary for Rotation`)
	sregex                   = regexp.MustCompile(`secondary: ([.[:alnum:]-_]+)$`)
)

// New creates a new Calendar.
func New(credFunc func(*router.Context) (*http.Client, error)) *Calendar {
	return &Calendar{
		credentials: credFunc,
	}
}

// CreateEvent creates new calendar events from the provided ShiftEntries.
func (c *Calendar) CreateEvent(ctx *router.Context, cfg *rotang.Configuration, shifts []rotang.ShiftEntry, sendUpdates bool) ([]rotang.ShiftEntry, error) {
	if err := ctx.Context.Err(); err != nil {
		return nil, err
	}
	client, err := c.credentials(ctx)
	if err != nil {
		return nil, err
	}
	cal, err := gcal.New(client)
	if err != nil {
		return nil, err
	}

	events, err := shiftsToEvents(ctx.Context, cfg, shifts)
	if err != nil {
		return nil, err
	}

	var resEvt []*gcal.Event
	for _, evt := range events {
		insert := cal.Events.Insert(cfg.Config.Calendar, evt)
		if sendUpdates {
			insert = insert.SendUpdates("all")
		}
		e, err := insert.Do()
		if err != nil {
			return nil, err
		}

		resEvt = append(resEvt, e)
	}

	return eventsToShifts(ctx.Context, &gcal.Events{
		Items: resEvt,
	}, cfg.Config.Name, &cfg.Config.Shifts)
}

// eventsRange specifies the time-range used when finding and matching Google calendar events to RotaNG shifts.
// The reason for searching instead of just making a query for the ID directly is that when querying for an ID'
// sometimes the resulting Calendar Event is returned with blank Start/End times.
const eventsRange = 2 * fullDay

// Event returns the information about the provided shift from the associated calendar event.
func (c *Calendar) Event(ctx *router.Context, cfg *rotang.Configuration, shift *rotang.ShiftEntry) (*rotang.ShiftEntry, error) {
	if err := ctx.Context.Err(); err != nil {
		return nil, err
	}
	// For some reason using the Google calendar Event.Get does not give Start and End times.
	shifts, err := c.Events(ctx, cfg, shift.StartTime.Add(-eventsRange), shift.EndTime.Add(eventsRange))
	if err != nil {
		return nil, err
	}
	return findShifts(ctx.Context, shifts, shift)
}

func findShifts(ctx context.Context, shifts []rotang.ShiftEntry, find *rotang.ShiftEntry) (*rotang.ShiftEntry, error) {
	for _, s := range shifts {
		if s.EvtID == find.EvtID {
			s.StartTime, s.EndTime = find.StartTime, find.EndTime
			return &s, nil
		}
		if s.StartTime == find.StartTime && s.EndTime == find.EndTime {
			logging.Infof(ctx, "shift: %v -> s: %v", find, s)
			return &s, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "calendar event not found")
}

// Events returns events from the specified time range.
func (c *Calendar) Events(ctx *router.Context, cfg *rotang.Configuration, from, to time.Time) ([]rotang.ShiftEntry, error) {
	if err := ctx.Context.Err(); err != nil {
		return nil, err
	}
	client, err := c.credentials(ctx)
	if err != nil {
		return nil, err
	}
	cal, err := gcal.New(client)
	if err != nil {
		return nil, err
	}

	events, err := cal.Events.List(cfg.Config.Calendar).
		ShowDeleted(false).SingleEvents(true).
		TimeMin(from.Format(time.RFC3339)).TimeMax(to.Format(time.RFC3339)).
		Q(cfg.Config.Name).
		OrderBy("startTime").Do()
	if err != nil {
		return nil, err
	}
	return eventsToShifts(ctx.Context, events, cfg.Config.Name, &cfg.Config.Shifts)
}

const fullDay = 24 * time.Hour

// TrooperShifts returns tropper oncallers for a timerange.
func (c *Calendar) TrooperShifts(ctx *router.Context, calendar, match, shiftName string, from, to time.Time) ([]rotang.ShiftEntry, error) {
	if err := ctx.Context.Err(); err != nil {
		return nil, err
	}
	client, err := c.credentials(ctx)
	if err != nil {
		return nil, err
	}
	cal, err := gcal.New(client)
	if err != nil {
		return nil, err
	}
	events, err := cal.Events.List(calendar).
		ShowDeleted(false).SingleEvents(true).
		TimeMin(from.Format(time.RFC3339)).TimeMax(to.Format(time.RFC3339)).
		Q(match).
		OrderBy("startTime").Do()
	if err != nil {
		return nil, err
	}
	return trooperToShifts(events, match, shiftName)
}

func trooperToShifts(events *gcal.Events, match, shiftName string) ([]rotang.ShiftEntry, error) {
	var res []rotang.ShiftEntry
	for _, e := range events.Items {
		start, err := calToTime(e.Start)
		if err != nil {
			return nil, err
		}
		end, err := calToTime(e.End)
		if err != nil {
			return nil, err
		}
		oc, err := oncallerFromSummary(e.Summary, match)
		if err != nil {
			return nil, err
		}
		var oncallers []rotang.ShiftMember
		for _, o := range oc {
			oncallers = append(oncallers, rotang.ShiftMember{
				Email:     o + "@google.com",
				ShiftName: shiftName,
			})
		}
		res = append(res, rotang.ShiftEntry{
			Name:      shiftName,
			StartTime: start,
			EndTime:   end,
			OnCall:    oncallers,
		})
	}
	return res, nil
}

// TrooperOncall returns trooper oncallers for specified time.
func (c *Calendar) TrooperOncall(ctx *router.Context, calendar, match string, t time.Time) ([]string, error) {
	if err := ctx.Context.Err(); err != nil {
		return nil, err
	}
	client, err := c.credentials(ctx)
	if err != nil {
		return nil, err
	}
	cal, err := gcal.New(client)
	if err != nil {
		return nil, err
	}

	events, err := cal.Events.List(calendar).
		ShowDeleted(false).SingleEvents(true).
		TimeMin(t.Add(-fullDay).Format(time.RFC3339)).TimeMax(t.Add(fullDay).Format(time.RFC3339)).
		Q(match).
		OrderBy("startTime").Do()
	if err != nil {
		return nil, err
	}

	return findTrooperEvent(events, match, t)
}

func findTrooperEvent(es *gcal.Events, match string, t time.Time) ([]string, error) {
	for _, e := range es.Items {
		start, err := calToTime(e.Start)
		if err != nil {
			return nil, err
		}
		end, err := calToTime(e.End)
		if err != nil {
			return nil, err
		}
		if (t.After(start) || t.Equal(start)) && t.Before(end) {
			return oncallerFromSummary(e.Summary, match)
		}
	}
	return nil, status.Errorf(codes.NotFound, "trooper oncall shift not found for time: %v", t)
}

func oncallerFromSummary(in, match string) ([]string, error) {
	var ocs []string

	// There are two different summary formats supported.

	// (1) "${primiary} is primiary for Rotation ${rotation_name} secondary: ${secondary}",
	// where
	// - ${rotation_name} must match with the value of `match`
	// - " secondary: ${secondary}" is optional, and
	if pri := pregex.FindStringSubmatch(in); pri != nil {
		// validate rotation name
		s := fmt.Sprintf("%s is primary for Rotation %s", pri[1], match)
		if !strings.HasPrefix(in, s) {
			return nil, status.Errorf(codes.InvalidArgument, "unmatched rotation: %q", in)
		}
		ocs = append(ocs, pri[1])
		if sec := sregex.FindStringSubmatch(in); sec != nil {
			ocs = append(ocs, sec[1])
		}
		return ocs, nil
	}

	// (2) "${prefix} ${primary}, ${secondary-1}, ${sec-2}, ... ${sec-N}",
	// where
	// - ${prefix} is the value of `match`, and
	// - all of ${primiary}, and ${secondary-N} are optional.
	if !strings.HasPrefix(in, match) {
		return nil, status.Errorf(codes.InvalidArgument, "event has invalid format")
	}
	tokens := strings.Split(strings.TrimPrefix(in, match), ",")
	for _, t := range tokens {
		ocs = append(ocs, strings.Join(strings.Fields(t), ""))
	}
	return ocs, nil
}

// nameShiftSeparator is used to separate the ShiftName from the rota name in Calendar Events.
const nameShiftSeparator = " ~s~ "

func eventsToShifts(ctx context.Context, events *gcal.Events, name string, shifts *rotang.ShiftConfig) ([]rotang.ShiftEntry, error) {
	if events == nil || shifts == nil || name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "all arguments must be set")
	}
	var res []rotang.ShiftEntry
	if len(shifts.Shifts) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "no shifts")
	}
	for _, e := range events.Items {
		if e.Id == "" {
			logging.Warningf(ctx, "Calendar event has empty ID, e: %v", e)
		}
		shift := shifts.Shifts[0].Name
		nm := strings.Split(e.Summary, nameShiftSeparator)
		if len(nm) > 1 {
			shift = nm[len(nm)-1]
		}
		// b/118075231, Q is not exact match.
		if nm[0] != name {
			continue
		}
		if e.Start == nil {
			logging.Warningf(ctx, "Start for e.Start: %v rota: %q is nil, evt: %v", e.Start, name, e)
			continue
		}
		start, err := calToTime(e.Start)
		if err != nil {
			return nil, err
		}
		if e.End == nil {
			logging.Warningf(ctx, "End for e.End: %v rota: %q is nil, evt: %v", e.End, name, e)
			continue
		}
		end, err := calToTime(e.End)
		if err != nil {
			return nil, err
		}
		var members []rotang.ShiftMember
		for _, a := range e.Attendees {
			if a.ResponseStatus == "declined" {
				continue
			}
			members = append(members, rotang.ShiftMember{
				Email:     a.Email,
				ShiftName: shift,
			})
		}
		res = append(res, rotang.ShiftEntry{
			Name:      shift,
			StartTime: start,
			EndTime:   end,
			OnCall:    members,
			Comment:   "Generated from calendar event",
			EvtID:     e.Id,
		})
	}
	return res, nil
}

func shiftsToEvents(ctx context.Context, cfg *rotang.Configuration, shifts []rotang.ShiftEntry) ([]*gcal.Event, error) {
	if cfg == nil {
		return nil, status.Errorf(codes.InvalidArgument, "all arguments must be set")
	}
	var res []*gcal.Event
	for _, s := range shifts {
		logging.Infof(ctx, "Reading Rota shift: %+v", s)
		var att []*gcal.EventAttendee
		sum := cfg.Config.Name
		if len(cfg.Config.Shifts.Shifts) > 1 {
			sum = cfg.Config.Name + nameShiftSeparator + s.Name
		}
		desc, err := fillCalendarDescription(cfg, s, sum)
		if err != nil {
			return nil, err
		}
		for _, m := range s.OnCall {
			att = append(att, &gcal.EventAttendee{
				Email: m.Email,
			})
		}
		start, end, err := timeToCal(cfg, s)
		if err != nil {
			return nil, err
		}
		evt := gcal.Event{
			Summary:     sum,
			Attendees:   att,
			Description: desc,
			Start:       start,
			End:         end,
			// Show calendar events as "Available" rather than "Busy".
			Transparency: "transparent",
		}
		logging.Infof(ctx, "Inserting Google Calendar event with START: %+v and END: %+v", start, end)
		res = append(res, &evt)
	}
	return res, nil
}

func fillCalendarDescription(cfg *rotang.Configuration, shift rotang.ShiftEntry, summary string) (string, error) {
	info := rotang.Info{
		RotaName:    summary,
		ShiftConfig: cfg.Config.Shifts,
		ShiftEntry:  shift,
	}
	descriptionTemplate, err := template.New("Description").Parse(cfg.Config.Description)
	if err != nil {
		return "", err
	}
	var descriptionBuf bytes.Buffer
	if err := descriptionTemplate.Execute(&descriptionBuf, &info); err != nil {
		return "", err
	}
	return descriptionBuf.String(), nil
}

const dayFormat = "2006-01-02"

var mtvTime = func() *time.Location {
	loc, err := time.LoadLocation("US/Pacific")
	if err != nil {
		panic(err)
	}
	return loc
}()

// timeToCal converts times to Google Calendar EventDateTimes.
// If the FullDay option is set the Date field of EvenDateTime will be used, if not DateTime will be used
// to set a full Start/End time of the shift.
func timeToCal(cfg *rotang.Configuration, shift rotang.ShiftEntry) (*gcal.EventDateTime, *gcal.EventDateTime, error) {
	if len(cfg.Config.Shifts.Shifts) > 1 && cfg.Config.Shifts.FullDayEvents {
		return nil, nil, status.Errorf(codes.InvalidArgument, "can't enable FullDayEvents with split shifts")
	}
	if cfg.Config.Shifts.FullDayEvents {
		return &gcal.EventDateTime{
				Date: shift.StartTime.Format(dayFormat),
			}, &gcal.EventDateTime{
				Date: shift.EndTime.Format(dayFormat),
			}, nil
	}

	return &gcal.EventDateTime{
			DateTime: shift.StartTime.Format(time.RFC3339),
		}, &gcal.EventDateTime{
			DateTime: shift.EndTime.Format(time.RFC3339),
		}, nil
}

func calToTime(calTime *gcal.EventDateTime) (time.Time, error) {
	tz := time.UTC
	if calTime.TimeZone != "" {
		var err error
		tz, err = time.LoadLocation(calTime.TimeZone)
		if err != nil {
			return time.Time{}, err
		}
	}
	if calTime.Date != "" {
		// This ends up being UTC default since no TZ is specified in the
		// legacy calendar events. Setting to MTV times makes things line
		// up better.
		if calTime.TimeZone == "" {
			return time.ParseInLocation(dayFormat, calTime.Date, mtvTime)
		}
		return time.ParseInLocation(dayFormat, calTime.Date, tz)
	}
	return time.ParseInLocation(time.RFC3339, calTime.DateTime, tz)
}

func keepDeclined(evt *calendar.Event) []*calendar.EventAttendee {
	var declined []*calendar.EventAttendee
	for _, attendee := range evt.Attendees {
		if attendee.ResponseStatus != "declined" {
			continue
		}
		declined = append(declined, attendee)
	}
	return declined
}

// UpdateEvent updates the calendar event with information from the provided updated shift.
func (c *Calendar) UpdateEvent(ctx *router.Context, cfg *rotang.Configuration, updated *rotang.ShiftEntry) (*rotang.ShiftEntry, error) {
	if err := ctx.Context.Err(); err != nil {
		return nil, err
	}

	client, err := c.credentials(ctx)
	if err != nil {
		return nil, err
	}
	cal, err := gcal.New(client)
	if err != nil {
		return nil, err
	}

	events, err := shiftsToEvents(ctx.Context, cfg, []rotang.ShiftEntry{*updated})
	if err != nil {
		return nil, err
	}

	if len(events) != 1 {
		return nil, status.Errorf(codes.NotFound, "wrong mumber of events returned")
	}

	currentEvent, err := cal.Events.Get(cfg.Config.Calendar, updated.EvtID).Do()
	if err != nil {
		return nil, err
	}

	events[0].Attendees = append(events[0].Attendees, keepDeclined(currentEvent)...)

	evt, err := cal.Events.Update(cfg.Config.Calendar, updated.EvtID, events[0]).Do()
	if err != nil {
		return nil, err
	}
	shifts, err := eventsToShifts(ctx.Context, &gcal.Events{
		Items: []*gcal.Event{evt},
	}, cfg.Config.Name, &cfg.Config.Shifts)
	if err != nil {
		return nil, err
	}

	if len(shifts) != 1 {
		return nil, status.Errorf(codes.NotFound, "wrong mumber of shifts returned")
	}

	res := shifts[0]
	return &res, err
}

// DeleteEvent deletes the calendar event matching the provided shift.
func (c *Calendar) DeleteEvent(ctx *router.Context, cfg *rotang.Configuration, shift *rotang.ShiftEntry) error {
	if err := ctx.Context.Err(); err != nil {
		return err
	}

	client, err := c.credentials(ctx)
	if err != nil {
		return err
	}
	cal, err := gcal.New(client)
	if err != nil {
		return err
	}

	return cal.Events.Delete(cfg.Config.Calendar, shift.EvtID).Do()
}

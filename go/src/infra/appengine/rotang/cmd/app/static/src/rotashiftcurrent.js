import {LitElement, html} from '@polymer/lit-element';
import {DateTime} from 'luxon';
import * as constants from './constants';

class RotaShiftCurrent extends LitElement {
  static get properties() {
    return {
      shifts: {type: constants.Shifts},
      updateState: {},
    };
  }

  constructor() {
    super();
  }

  addMember(splitIdx, shiftIdx) {
    if (!this.shifts.SplitShifts[splitIdx].Shifts[shiftIdx].OnCall) {
      this.shifts.SplitShifts[splitIdx].Shifts[shiftIdx].OnCall = [];
    }
    let shift = this.shifts.SplitShifts[splitIdx].Shifts[shiftIdx];
    shift.OnCall = [...shift.OnCall, {
      Email: this.shifts.SplitShifts[splitIdx].Members[0],
      ShiftName: this.shifts.SplitShifts[splitIdx].Name,
    }];
    this.requestUpdate();
  }

  removeMember(splitIdx, shiftIdx, oncallIdx) {
    let shift = this.shifts.SplitShifts[splitIdx].Shifts[shiftIdx];
    shift.OnCall = shift.OnCall.slice(0, oncallIdx)
      .concat(shift.OnCall.slice(oncallIdx+1));
    this.requestUpdate();
  }

  removeShifts(splitIdx, shiftIdx) {
    this.shifts.SplitShifts[splitIdx].Shifts =
      this.shifts.SplitShifts[splitIdx].Shifts.slice(0, shiftIdx);
    this.requestUpdate();
  }

  fixComments() {
    for (let i = 0; i < this.shifts.SplitShifts.length; i++) {
      for (let j = 0; j < this.shifts.SplitShifts[i].Shifts.length; j++) {
        this.shifts.SplitShifts[i].Shifts[j].Comment =
          this.shadowRoot.getElementById(this.commentID(i, j)).value;
      }
    }
  }

  commentID(splitIDX, shiftIDX) {
    return `shiftComment_${splitIDX}-${shiftIDX}`;
  }

  shiftMemberID(splitIdx, shiftIdx, oncallIdx) {
    return `shiftMember_${splitIdx}_${shiftIdx}_${oncallIdx}`;
  }

  async sendJSON() {
    this.fixComments();
    try {
      const res = await fetch('shiftsupdate', {
        method: 'POST',
        body: JSON.stringify(this.shifts),
        headers: {
          'Content-Type': 'application/json',
        },
      });
      if (!res.ok) {
        throw res;
      }
      this.updateState = html`<small class="ok">(Shifts updated)</small>`;
    } catch (err) {
      this.updateState = html`<small class="fail">${err.text()}</small>`;
    }
  }

  setMember(splitIdx, shiftIdx, oncallIdx) {
    this.shifts.SplitShifts[splitIdx].Shifts[shiftIdx].OnCall[oncallIdx].Email =
      this.shadowRoot.getElementById(
        this.shiftMemberID(splitIdx, shiftIdx, oncallIdx)).value;
  }

  selectOnCallers(splitIdx, shiftIdx, shift, members) {
    if (!shift.OnCall) {
      return;
    }
    return shift.OnCall.map((oncall, oncallIdx) => html`
    <select id="${this.shiftMemberID(splitIdx, shiftIdx, oncallIdx)}"
      @change=${() => this.setMember(splitIdx, shiftIdx, oncallIdx)}>
      <option value="${oncall.Email}">${oncall.Email}</option>
      ${members.map((o) => html`<option value=${o}>${o}</option>`)}
    </select>
    <button type="button" @click=${() =>
    this.removeMember(splitIdx, shiftIdx, oncallIdx)}>
      <small>-</small>
    </button>
    `);
  }

  currentShifts() {
    if (!this.shifts || !this.shifts.SplitShifts) {
      return;
    }
    return this.shifts.SplitShifts.map((s, splitIdx) => html`
    <h4>${s.Name}</h4>
      <table id="shifts">
      <thead>
        <th>Oncallers</th>
        <th>Start</th>
        <th>End</th>
       <th>Comment</th>
      <thead>
      <tbody>
        ${this.shiftTemplate(s, splitIdx)}
      </tbody>
      `);
  }

  shiftTemplate(ss, splitIdx) {
    return ss.Shifts.map((i, shiftIdx) => html`
      <tr>
        <td>
          ${this.selectOnCallers(splitIdx, shiftIdx, i, ss.Members)}
          <button type="button" @click=
            ${() => this.addMember(splitIdx, shiftIdx)}>
        <small>+</small>
        </button>
        </td>
        <td>
          <input type="text" name="shiftStart" class="roInput"
        value="${DateTime.fromISO(i.StartTime, {zone: constants.zone}).toFormat(constants.timeFormat)}" readonly>
        </td>
        <td>
          <input type="text" name="shiftEnd" class="roInput"
        value="${DateTime.fromISO(i.EndTime, {zone: constants.zone}).toFormat(constants.timeFormat)}" readonly>
        </td>
        <td>
          <input type=text id="${this.commentID(splitIdx, shiftIdx)}"
          value="${i.Comment}">
        </td>
        <td>
          <div class="tooltip">
            <span class="tooltiptext">
              Clear all shifts from this onwards.
            </span>
            <button type="button" @click=
                  ${() => this.removeShifts(splitIdx, shiftIdx)}>&#x21e3;
            </button>
          </div>
        </td>
      </tr>`);
  }

  render() {
    return html`
      <style>
        .tooltip {
          position: relative;
          display: inline-block;
          border-bottom: 1px dotted black;
        }

        .tooltip .tooltiptext {
          visibility: hidden;
          width: 240px;
          bottom: 100%;
          left: 50%;
          mergin-left: -120px;
          background-color: black;
          color: #fff;
          text-align: center;
          padding: 5px 0;
          border-radius: 6px;
          position: absolute;
          z-index: 1;
        }

        .tooltip:hover .tooltiptext {
          visibility: visible;
        }

        #shifts {
          font-size: small;
          border-collapse: collapse;
        }

        #shifts td, #shifts th {
          border: 1px, solid #ddd;
          padding: 8px;
        }

        #shifts th {
          text-align: left;
        }

        .roInput {
          font-size: small;
          border: 0px solid;
          background-color: transparent;
        }

        .ok {
          color: green;
          background-color: transparent;
        }

        .fail {
          color: red;
          background-color: transparent;
        }

        #shifts tr:nth-child(even) {
          background-color: hsl(0, 0%, 95%);
        };
      </style>
      <form action="modifyshifts" method=POST>
        <fieldset>
          <legend>Current shifts</legend>
          ${this.currentShifts()}
            <br>
              <button type="button" @click=${() => this.sendJSON()}>
              Update Shifts</button>
              <small>${this.updateState && this.updateState}</small>
          </table>
        </fieldset>
      </form>
      `;
  }
}

customElements.define('rota-shift-current', RotaShiftCurrent);

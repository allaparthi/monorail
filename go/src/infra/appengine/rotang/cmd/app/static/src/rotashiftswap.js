import {LitElement, html} from '@polymer/lit-element';
import {DateTime} from 'luxon';
import * as constants from './constants';

class RotaShiftSwap extends LitElement {
  static get properties() {
    return {
      shifts: {type: constants.Shifts},
      neededMembers: {},
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
      const res = await fetch('/shiftswap', {
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
      ${this.userOnCall(splitIdx, shiftIdx, oncallIdx, members, oncall)}
    `);
  }

  userOnCall(splitIdx, shiftIdx, oncallIdx, members, oncall) {
    if (!members) {
      return html`${oncall.Email}`;
    }
    if (members[0] === oncall.Email) {
      return html`<select name="shiftMember">
        <option value="${oncall.Email}">${oncall.Email}</option>
      </select>
        ${this.deleteButton(splitIdx, shiftIdx, oncallIdx)}
      `;
    }
    return html`<select
      id="${this.shiftMemberID(splitIdx, shiftIdx, oncallIdx)}"
      @change=${() => this.setMember(splitIdx, shiftIdx, oncallIdx)}>
      <option value="${oncall.Email}">${oncall.Email}</option>
      <option value=${members[0]}>${members[0]}</option>
    </select>`;
  }

  deleteButton(splitIdx, shiftIdx, oncallIdx) {
    if (this.neededMembers != 0 &&
      this.shifts.SplitShifts[splitIdx].Shifts[shiftIdx].OnCall.length <=
      this.neededMembers) {
      return;
    }
    return html`
      <button type="button"
        @click=${() => this.removeMember(splitIdx, shiftIdx, oncallIdx)}>
        <small>-</small>
      </button>
    `;
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
      </tr>`);
  }

  render() {
    return html`
      <style>
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
      <form action="swapshifts" method=POST>
        <fieldset>
          <legend>Current shifts</legend>
          ${this.currentShifts()}
            <br>
              <button type="button" @click=${() => this.sendJSON()}>
              Swap Shifts</button>
              <small>${this.updateState && this.updateState}</small>
          </table>
        </fieldset>
      </form>
      `;
  }
}

customElements.define('rota-shift-swap', RotaShiftSwap);

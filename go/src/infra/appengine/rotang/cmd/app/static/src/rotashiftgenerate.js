import {LitElement, html} from '@polymer/lit-element';
import {DateTime} from 'luxon';
import * as constants from './constants';

class RotaShiftGenerate extends LitElement {
  static get properties() {
    return {
      shifts: {type: constants.Shifts},
      rota: {type: String},
      generators: {type: constants.Shifts},
      modifiers: {type: constants.Shifts},
      updateState: {},
    };
  }

  constructor() {
    super();
    this.generated = false;
  }

  async generate() {
    const startElement = this.shadowRoot.getElementById('start').value;
    const numberElement = this.shadowRoot.getElementById('nrToSchedule').value;
    const generatorElement = this.shadowRoot.getElementById('generator').value;
    let formBody = new URLSearchParams();
    formBody.append('name', this.rota);
    formBody.append('nrShifts', numberElement);
    formBody.append('startTime', startElement);
    formBody.append('generator', generatorElement);
    formBody.append('modifiers', this.buildModifiers());
    try {
      const res = await fetch('generate', {
        method: 'POST',
        body: formBody,
        headers: new Headers({
          'Content-type': 'application/x-www-form-urlencoded; charset=UTF-8',
        }),
      });
      this.shifts = await res.json();
    } catch (err) {
      console.log(err);
      this.updateState = html`<small class="fail">${err}</small>`;
    }

    this.generated = true;
  }

  async sendJSON() {
    try {
      const res = await fetch('shiftsgenerate', {
        method: 'POST',
        body: JSON.stringify(this.shifts),
        headers: {
          'Content-Type': 'application/json',
        },
      });
      if (!res.ok) {
        throw res;
      }
      window.location.reload(true);
      this.updateState = html`<small class="ok">OK</small>`;
    } catch (err) {
      this.updateState = html`<small class="fail">${err}</small>`;
    }
  }

  shiftMemberID(splitIdx, shiftIdx, oncallIdx) {
    return `shiftMember_${splitIdx}_${shiftIdx}_${oncallIdx}`;
  }

  setMember(splitIdx, shiftIdx, oncallIdx) {
    this.shifts.SplitShifts[splitIdx].Shifts[shiftIdx].OnCall[oncallIdx].Email =
      this.shadowRoot.getElementById(
        this.shiftMemberID(splitIdx, shiftIdx, oncallIdx)).value;
  }

  selectOnCallers(splitIdx, shiftIdx, ss, shift) {
    if (!shift.OnCall) {
      return;
    }
    return shift.OnCall.map((oncall, oncallIdx) => html`
     <select id="${this.shiftMemberID(splitIdx, shiftIdx, oncallIdx)}"
        @change=${() => this.setMember(splitIdx, shiftIdx, oncallIdx)})>
       <option value="${oncall.Email}">${oncall.Email}</option>
       ${ss.Members.map((o) => html`<option value=${o}>${o}</option>`)}
     </select>`);
  }

  renderShifts() {
    if (!this.shifts || !this.shifts.SplitShifts) {
      return;
    }
    return this.shifts.SplitShifts.map((s, splitIdx) => html`
      <h4>${s.Name}</h4>
      <table class="shifts">
      <thead>
        <th>Oncallers</th>
        <th>Start</th>
        <th>End</th>
        <th>Comment</th>
      <thead>
      <tbody>
        ${this.shiftsTemplate(splitIdx, s)}
      </tbody>
      </table>`);
  }

  fillGenerators() {
    if (!this.generators) {
      return;
    }
    return html`
      <select name="generator" id="generator">
        ${this.generators.map((g) => html`
          <option value=${g}>${g}</option>
        `)};
      </select>
    `;
  }

  fillModifiers() {
    if (!this.modifiers) {
      return;
    }
    const res = this.modifiers.map((mod, i) => {
      const name = html`<small>${mod.Name}</small><br>`;
      if (mod.Checked) {
        return html`<input type="checkbox" name="${mod.Name}"
          value="${mod.Name}" id="modifier-${i}"
          checked>${name}`;
      }
      return html`<input type="checkbox" name="${mod.Name}" value="${mod.Name}"
        id="modifier-${i}">${name}`;
    });
    return html`<span class="flex-container">${res}</span>`;
  }

  buildModifiers() {
    const modifiers = [];
    if (!this.modifiers) {
      return modifiers;
    }
    for (let i = 0; i < this.modifiers.length; i++) {
      if (this.shadowRoot.getElementById(`modifier-${i}`).checked) {
        modifiers.push(this.shadowRoot.getElementById(`modifier-${i}`).value);
      }
    }
    return modifiers;
  }

  shiftsTemplate(splitIdx, ss) {
    return ss.Shifts.map((i, shiftIdx) => html`
      <tr>
        <td>
        <table>
        <tbody>
        <td>
           ${this.selectOnCallers(splitIdx, shiftIdx, ss, i)}
        </td>
        </tbody>
        </table>
        </td>
        <td>${DateTime.fromISO(i.StartTime, {zone: constants.zone}).toFormat(constants.timeFormat)}</td>
        <td>${DateTime.fromISO(i.EndTime, {zone: constants.zone}).toFormat(constants.timeFormat)}</td>
        <td>${i.Comment}</td>
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

        .shifts {
          font-size: small;
          border-collapse: collapse;
        }

        .shifts td, .shifts th {
          border: 1px, solid #ddd;
          padding: 8px;
        }

        .shifts th {
          text-align: left;
        }

        .ok {
          color: green;
          background-color: transparent;
        }

        .fail {
          color: red;
          background-color: transparent;
        }

        .shifts tr:nth-child(even) {
          background-color: hsl(0, 0%, 95%);
        };
        .flex-container {
          display: flex;
        }
      </style>
      <fieldset>
      <form>
        <table>
          <tbody>
            <tr>
              <td>
                <div class="tooltip">
                  <span class="tooltiptext">
                    Leaving this empty will generate from the end of last
                    current shift.
                  </span>
                  StartDate:<input type="date" id="start"></td>
                </div>
              <td>
                <div class="tooltip">
                  <span class="tooltiptext">
                    Leaving this empty will use the configuration default.
                  </span>
                  ShiftsToSchedule:<input type="number" id="nrToSchedule"></td>
                </div>
              <td>Generator:
                ${this.fillGenerators()}
              </td>
              <td>Modifiers:
                ${this.fillModifiers()}
              </td>
            </tr>
          </tbody>
        </table>
      </form>
      <button type="button" @click=${() => this.generate()}>Generate</button>
      ${this.generated ? html`
        <button type="button" @click=${() => this.sendJSON()}>Submit</button>
        <small>${this.updateState && this.updateState}</small>`: html``}
      <form>
        ${this.renderShifts()}
      </form>
      </fieldset>
      `;
  }
}

customElements.define('rota-shift-generate', RotaShiftGenerate);

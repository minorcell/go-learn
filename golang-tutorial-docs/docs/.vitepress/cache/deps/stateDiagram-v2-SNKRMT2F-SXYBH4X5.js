import {
  StateDB,
  stateDiagram_default,
  stateRenderer_v3_unified_default,
  styles_default
} from "./chunk-ONO2RB76.js";
import "./chunk-4I4ID2TI.js";
import "./chunk-PJ3UDJCS.js";
import "./chunk-2B5L2MTA.js";
import "./chunk-2QJ2RYLZ.js";
import "./chunk-TB3MTNXB.js";
import "./chunk-3PDFNL3P.js";
import "./chunk-WOJRECXA.js";
import "./chunk-TQHQVS5X.js";
import "./chunk-LQRCPGYC.js";
import {
  __name
} from "./chunk-TTPLXKQF.js";
import "./chunk-4UTD2NOI.js";
import "./chunk-FDBJFBLO.js";

// node_modules/mermaid/dist/chunks/mermaid.core/stateDiagram-v2-SNKRMT2F.mjs
var diagram = {
  parser: stateDiagram_default,
  get db() {
    return new StateDB(2);
  },
  renderer: stateRenderer_v3_unified_default,
  styles: styles_default,
  init: __name((cnf) => {
    if (!cnf.state) {
      cnf.state = {};
    }
    cnf.state.arrowMarkerAbsolute = cnf.arrowMarkerAbsolute;
  }, "init")
};
export {
  diagram
};
//# sourceMappingURL=stateDiagram-v2-SNKRMT2F-SXYBH4X5.js.map

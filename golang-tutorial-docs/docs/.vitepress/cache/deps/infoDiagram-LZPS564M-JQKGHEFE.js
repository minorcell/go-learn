import {
  parse
} from "./chunk-VC7HE666.js";
import "./chunk-BGQACCT4.js";
import "./chunk-GEEEUXMY.js";
import "./chunk-IYCHR5BC.js";
import "./chunk-5ABCPFVG.js";
import "./chunk-JSD5TSMO.js";
import "./chunk-NS2P4TKQ.js";
import "./chunk-32BEVRBU.js";
import "./chunk-W4C6O4J6.js";
import {
  package_default
} from "./chunk-OPB6PG5C.js";
import {
  selectSvgElement
} from "./chunk-WHKN7ZWP.js";
import {
  __name,
  configureSvgSize,
  log
} from "./chunk-TTPLXKQF.js";
import "./chunk-2VRVB2MD.js";
import "./chunk-4UTD2NOI.js";
import "./chunk-FDBJFBLO.js";

// node_modules/mermaid/dist/chunks/mermaid.core/infoDiagram-LZPS564M.mjs
var parser = {
  parse: __name(async (input) => {
    const ast = await parse("info", input);
    log.debug(ast);
  }, "parse")
};
var DEFAULT_INFO_DB = {
  version: package_default.version + (true ? "" : "-tiny")
};
var getVersion = __name(() => DEFAULT_INFO_DB.version, "getVersion");
var db = {
  getVersion
};
var draw = __name((text, id, version) => {
  log.debug("rendering info diagram\n" + text);
  const svg = selectSvgElement(id);
  configureSvgSize(svg, 100, 400, true);
  const group = svg.append("g");
  group.append("text").attr("x", 100).attr("y", 40).attr("class", "version").attr("font-size", 32).style("text-anchor", "middle").text(`v${version}`);
}, "draw");
var renderer = { draw };
var diagram = {
  parser,
  db,
  renderer
};
export {
  diagram
};
//# sourceMappingURL=infoDiagram-LZPS564M-JQKGHEFE.js.map

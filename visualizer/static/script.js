var s;
var data;
$(document).ready(function () {
  $.getJSON("./static/data.json", function (data) {
    s = new sigma({
      graph: data,
      container: "container",
      settings: {
        defaultNodeColor: "#ec5148",
        defaultLabelColor: "#ffffff",
        edgeColor: "target",
      },
    });
    data.paths.forEach((path) => {
      for (var i = 0; i < path.ants; i++) {
        path.nodes.forEach((room) => {
          s.graph.nodes().forEach((node) => {
            if (node.id == room) {
              node.color = "#ffffff";
              s.refresh();
              setTimeout(() => {}, 100);
              node.color = "default";
              s.refresh();
            }
          });
        });
      }
    });
    console.log(s.interNodes());
  });
});

function getEdgeID(source, target) {
  var id;
  var edges = s.graph.edges();
  for (var i = 0; i < edges.length; i++) {
    if ((edges[i].source == source && edges[i].target == target) || (edges[i].source == target && edges[i].target == source)) {
      id = edges[i].id;
      break;
    }
  }
  return id;
}

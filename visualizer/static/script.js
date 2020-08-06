var cy;
var jsonData;
var antsAmount;
$(document).ready(function () {
  $.getJSON("./static/data.json", function (data) {
    jsonData = data;
    initCy();
    data.paths.forEach(function (path) {
      var color = "#" + Math.floor(Math.random() * 16777215).toString(16);
      path.nodes.forEach((node, index) => {
        if (node != cy.nodes('node[type = "start"]').id() && node != cy.nodes('node[type = "end"]').id()) {
          cy.nodes('node[id = "' + node + '"]').style({
            "background-color": color,
            opacity: 0.3,
          });
        }
        if (index > 0 && index < path.nodes.length) {
          cy.$id(path.edges[index - 1]).style({
            "line-color": color,
          });
        }
      });
    });
    $("#play").click(function (e) {
      e.preventDefault();
      jsonData.paths.forEach((path) => {
        path.nodes.forEach((node, index) => {
          for (var i = 0; i < path.ants; i++) {
            setTimeout(function () {
              if (index > 0 && index < node.length) {
                cy.$id(path.edges[index - 1])
                  .animate({ style: { opacity: 1 } }, { duration: 20 })
                  .delay(400)
                  .animate({ style: { opacity: 0.2 } });
              }
              cy.$id(node)
                .animate({ style: { opacity: 1 } }, { duration: 20 }, { queue: true })
                .delay(400)
                .animate({ style: { opacity: 0.3 } });
            }, 1000 * index);
          }
        });
      });
    });
    $("#reset").click(function (e) {
      e.preventDefault();
      cy.nodes().stop();
      cy.edges().stop();
      initCy;
    });
  });
});

function initCy() {
  cy = cytoscape({
    container: $("#cy"),
    elements: jsonData,
    layout: { name: "cose-bilkent" },
    style: [
      {
        selector: "node",
        style: {
          content: "data(id)",
          opacity: "0.05",
          color: "#fff",
          "font-size": "10px",
          "font-family": "Monaco",
          "text-valign": "center",
          "text-halign": "center",
          "background-color": "#000000",
          "z-index": "10",
        },
      },
      {
        selector: "edge",
        style: {
          "overlay-padding": "3px",
          "curve-style": "straight",
          opacity: "0.2",
        },
      },
    ],
  });
  cy.nodes('node[type = "start"]').style({
    "background-color": "#ffffff",
    "border-style": "solid",
    "border-color": "#00ff00",
    "border-width": "3px",
    color: "#000000",
    opacity: "1",
  });
  cy.nodes('node[type = "end"]').style({
    "background-color": "#ffffff",
    "border-style": "solid",
    "border-color": "#ff0000",
    "border-width": "3px",
    color: "#000000",
    opacity: "1",
  });
}

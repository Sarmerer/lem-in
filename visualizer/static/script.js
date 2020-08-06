var cy;
var jsonData;
var startNode;
var endNode;
var antsOut = 0;
$(document).ready(function () {
  $.getJSON("./static/data.json", function (data) {
    jsonData = data;
    initCy();
    startNode = data.paths[0].nodes[0];
    endNode = data.paths[0].nodes[data.paths[0].nodes.length - 1];
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
      var counter = 0;
      jsonData.paths.forEach((path, pIndex) => {
        moveAnts(path, pIndex);
        jsonData.paths.forEach((path) => {
          path.nodes.forEach((node) => {
            cy.$id(node).animate({ style: { opacity: 0.3 } }, { duration: 20 }, { queue: true });
          });
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

function moveAnts(path, pIndex) {
  while (path.antsInNodes[path.nodes.length - 1] != path.antsInPath) {
    for (var nIndex = 0; nIndex < path.nodes.length; nIndex++) {
      var node = path.nodes[nIndex];
      var nextNode = jsonData.paths[pIndex].antsInNodes[nIndex + 1];
      var currNode = jsonData.paths[pIndex].antsInNodes[nIndex];
      if (node != endNode && path.nodes[nIndex + 1] == endNode) {
        jsonData.paths[pIndex].antsInNodes[nIndex + 1] += 1;
        jsonData.paths[pIndex].antsInNodes[nIndex] = 0;
        color(node, currNode);
        antsOut++;
        break;
      } else if (currNode >= 1 && nextNode == 0) {
        jsonData.paths[pIndex].antsInNodes[nIndex]--;
        jsonData.paths[pIndex].antsInNodes[nIndex + 1] = 1;
        color(node, currNode);
        antsOut++;
        break;
      }
      $("#ants").text(antsOut + "/" + jsonData.ants);
    }
    // setTimeout(function () {
    //   if (index > 0 && index < node.length) {
    //     cy.$id(path.edges[index - 1])
    //       .animate({ style: { opacity: 1 } }, { duration: 20 })
    //       .delay(400)
    //       .animate({ style: { opacity: 0.2 } });
    //   }
    //   cy.$id(node)
    //     .animate({ style: { opacity: 1 } }, { duration: 20 }, { queue: true })
    //     .delay(400)
    //     .animate({ style: { opacity: 0.3 } });
    // }, 1000 * index);
    console.log(path.antsInNodes);
  }
}

function color(node, currNode) {
  if (currNode > 0 && node != startNode) {
    cy.$id(node).animate({ style: { opacity: 1 } }, { duration: 200 }, { queue: true });
  } else if (node == 0 && node != startNode) {
    cy.$id(node).animate({ style: { opacity: 0.3 } }, { duration: 200 }, { queue: true });
  }
}

function initCy() {
  $("#ants").text(antsOut + "/" + jsonData.ants);
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

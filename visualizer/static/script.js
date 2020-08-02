$(document).ready(function () {
  $.getJSON("./static/data.json", function (jsonData, textStatus, jqXHR) {
    console.log(jsonData);

    // create a network
    var container = document.getElementById("mynetwork");

    var options = {
      nodes: {
        font: {
          size: 20,
        },
        borderWidth: 3,
      },
      edges: {
        width: 5,
        color: "rgba(154, 18, 179, 0.3)",
      },
    };

    // initialize your network!
    var network = new vis.Network(container, jsonData, options);
    console.log(network);
    jsonData.paths.forEach((element) => {
      element.nodes.forEach((node) => {
        console.log(network.findNode(node));
      });
    });
  });
});

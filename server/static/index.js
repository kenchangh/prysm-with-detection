var intervalID;
var validators = ["1", "2"];
var seenAttests = [];

$(".start-btn").click(function() {
  intervalID = startLoadingData();
});

$(".pause-btn").click(function() {
  clearInterval(intervalID);
});

function startLoadingData() {
  return setInterval(function() {
    jQuery.get("http://localhost:5000/validator", function(data) {
      console.log(data);
      addNewAttestations(data);
    });
  }, 1000);
}

function addNewAttestations(data) {
  const keys = Object.keys(data.selected);
  const newKeys = keys.filter(key => {
    const notSeen = seenAttests.indexOf(key) == -1;
    return notSeen;
  });

  newKeys.forEach(key => {
    seenAttests.push(key);
  });

  newKeys.forEach(key => {
    const val1color = "green";
    let val2color = "green";
    let resultColor = "green";

    const val1Attest = data.validators["1"][key];
    const val2Attest = data.validators["2"][key];
    const selectedAttest = data.selected[key];

    if (val2Attest.dataStr !== val1Attest.dataStr) {
      val2color = "red";

      if (selectedAttest.dataStr === val2Attest.dataStr) {
        resultColor = "red";
      }
    }

    const results = [val1color, val2color, resultColor];
    renderBoxes(results);
  });
}

function renderBoxes(colors) {
  var boxes = colors.map(function(color) {
    return `<div class="box ${color}-box"></div>`;
  });

  document.getElementById("data").innerHTML += `
    <div class="row">
      ${boxes.join("")}
    </div>
  `;
}

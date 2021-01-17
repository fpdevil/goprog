// Get the svg element for drawing
var qSvgTag = document.getElementById("WirePlotSvg");

function Point3D(i, j, daaZ) {
  this.mdX = 10 * i - 200;
  this.mdY = 10 * j - 200;
  this.mdZ = daaZ[i][j];
}

function CProjectedPoint(dX, dY, dZ) {
  var dScale = 1.5;
  var dDuDx = 0.612;
  var dDuDy = -dDuDx;
  var dDuDz = 0.0;
  var dDvDx = 0.25;
  var dDvDy = 0.25;
  var dDvDz = -0.866;
  this.mdU = (dDuDx * dX + dDuDy * dY + dDuDz * dZ) * dScale + 400;
  this.mdV = (dDvDx * dX + dDvDy * dY + dDvDz * dZ) * dScale + 400;
}

var daaZ = new Array(41);
for (var i = 0; i < 41; ++i) {
  daaZ[i] = new Array(41);
  for (var j = 0; j < 41; ++j) {
    var dX = 10 * i - 200;
    var dY = 10 * j - 200;
    daaZ[i][j] = 300 * Math.exp(-(dX * dX + dY * dY) / 10000);
  }
}

for (var i = 0; i < 40; ++i) {
  for (var j = 0; j < 40; ++j) {
    var qP1 = new Point3D(i, j, daaZ);
    var qP2 = new Point3D(i + 1, j, daaZ);
    var qP3 = new Point3D(i, j + 1, daaZ);
    var qP4 = new Point3D(i + 1, j + 1, daaZ);

    var qProj1 = new CProjectedPoint(qP1.mdX, qP1.mdY, qP1.mdZ);
    var qProj2 = new CProjectedPoint(qP2.mdX, qP2.mdY, qP2.mdZ);
    var qProj3 = new CProjectedPoint(qP3.mdX, qP3.mdY, qP3.mdZ);
    var qProj4 = new CProjectedPoint(qP4.mdX, qP4.mdY, qP4.mdZ);
    var dCoords = qProj1.mdU +
    " " +
    qProj1.mdV +
    " " +
    qProj2.mdU +
    " " +
    qProj2.mdV +
    " " +
    qProj4.mdU +
    " " +
    qProj4.mdV +
    " " +
    qProj3.mdU +
    " " +
    qProj3.mdV;

    var qPolygon = document.createElementNS(
      "http://www.w3.org/2000/svg",
      "polygon"
    );
    qPolygon.setAttribute("points", dCoords);
    qPolygon.setAttribute("fill", "#ffffff");
    qPolygon.setAttribute("fill-rule", "nonzero");
    qPolygon.setAttribute("stroke-width", "1");
    qPolygon.setAttribute("stroke", "rgb(0, 128, 0)");
    qSvgTag.appendChild(qPolygon);
  }
}

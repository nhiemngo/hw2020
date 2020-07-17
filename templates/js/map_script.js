var lat = document.getElementById("generate-map").getAttribute("lat");
var lng = document.getElementById("generate-map").getAttribute("lng");

var name = document.getElementById("generate-map").getAttribute("name");

var iconFeature = new ol.Feature({
  geometry: new ol.geom.Point(ol.proj.transform([lng, lat], 'EPSG:4326', 'EPSG:3857')),
  name: name,
  population: 4000,
  rainfall: 500
});

var iconStyle = new ol.style.Style({
  image: new ol.style.Icon( /** @type {olx.style.IconOptions} */ ({
    anchor: [0.5, 46],
    anchorXUnits: 'fraction',
    anchorYUnits: 'pixels',
    src: 'https://openlayers.org/en/v4.6.5/examples/data/icon.png'
  }))
});

iconFeature.setStyle(iconStyle);

var vectorSource = new ol.source.Vector({
  features: [iconFeature]
});

var vectorLayer = new ol.layer.Vector({
  source: vectorSource
});

var rasterLayer = new ol.layer.Tile({
  source: new ol.source.OSM()
});

var map = new ol.Map({
  layers: [rasterLayer, vectorLayer],
  target: 'map',
  controls: ol.control.defaults({
    attributionOptions: {
      collapsible: false
    }
  }),
  view: new ol.View({
    center: ol.proj.fromLonLat([lng, lat]),
    zoom: 10
  })
});

var element = document.getElementById('popup');

var popup = new ol.Overlay({
  element: element,
  positioning: 'bottom-center',
  stopEvent: false,
  offset: [0, -50]
});

map.addOverlay(popup);

map.on('pointermove', function(evt) {
  var feature = map.forEachFeatureAtPixel(evt.pixel,
      function(feature) {
        return feature;
      });
  if (feature) {
    var coordinates = feature.getGeometry().getCoordinates();
    popup.setPosition(coordinates);
    $(element).popover({
      placement: 'top',
      html: true,
      content: feature.get('name')
    });
    $(element).popover('show');
  } else {
    $(element).popover('dispose');
  }
});

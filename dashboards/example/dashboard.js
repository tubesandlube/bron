var blessed = require('blessed')
  , contrib = require('../../index')

var screen = blessed.screen()

// XXX
// authors and number commits they've contributed (all time)
// number of lines per language (all time)
// number of languages (timeline)
// number of lines (timeline)
// number of commits (timeline)
// number of files (timeline)

//create layout and widgets

var grid = new contrib.grid({rows: 1, cols: 2})

var grid1 = new contrib.grid({rows: 1, cols: 1})
grid1.set(0, 0, contrib.log,
  { fg: "green"
  , selectedFg: "green"
  , label: 'Number of Files'})

var grid3 = new contrib.grid({rows: 2, cols: 1})
grid3.set(0, 0, contrib.bar,
  { label: 'Lines per Language'
  , barWidth: 3
  , barSpacing: 6
  , xOffset: 0
  , maxHeight: 5})
grid3.set(1, 0, contrib.table,
  { keys: true
  , fg: 'green'
  , label: 'Number of Commits per Author'
  , columnSpacing: [24, 10]})

var grid4 = new contrib.grid({rows: 3, cols: 1})
grid4.set(0, 0, contrib.line,
  { style: 
    { line: "red"
    , text: "white"
    , baseline: "black"}
  , label: 'Number of Lines'
  , maxY: 60})
grid4.set(1, 0, grid3)
grid4.set(2, 0, grid1)

var grid5 = new contrib.grid({rows: 2, cols: 1})
grid5.set(0, 0, contrib.line,
  { showNthLabel: 5
  , maxY: 100
  , label: 'Number of Languages'})
grid5.set(1, 0, contrib.map, {label: 'Number of Commits'})
grid.set(0, 0, grid5)
grid.set(0, 1, grid4)

grid.applyLayout(screen)

var transactionsLine = grid5.get(0, 0)
var errorsLine = grid4.get(0, 0)
var map = grid5.get(1, 0)
var log = grid1.get(0, 0)
var table = grid3.get(1,0)
var bar = grid3.get(0, 0)

//dummy data
var servers = ['foo', 'bar', 'lalla', 'US1', 'US2', 'EU1', 'AU1', 'AS1', 'JP1']
var commands = ['grep', 'node', 'java', 'timer', '~/ls -l', 'netns', 'watchdog', 'gulp', 'tar -xvf', 'awk', 'npm install']

//set dummy data on bar chart
function fillBar() {
  var arr = []
  for (var i=0; i<servers.length; i++) {
    arr.push(Math.round(Math.random()*10))
  }
  bar.setData({titles: servers, data: arr})
}
fillBar()

//set log dummy data
var rnd = Math.round(Math.random()*2)
if (rnd==0) log.log('starting process ' + commands[Math.round(Math.random()*(commands.length-1))])
else if (rnd==1) log.log('terminating server ' + servers[Math.round(Math.random()*(servers.length-1))])
else if (rnd==2) log.log('avg. wait time ' + Math.random().toFixed(2))

//set dummy data for table
function generateTable() {
   var data = []

   for (var i=0; i<30; i++) {
     var row = []
     row.push(commands[Math.round(Math.random()*(commands.length-1))])
     row.push(Math.round(Math.random()*5))

     data.push(row)
   }

   table.setData({headers: ['Author', 'Commits'], data: data})
}

generateTable()
table.focus()

//set map dummy markers
var marker = true
setInterval(function() {
   if (marker) {
    map.addMarker({"lon" : "37.5000", "lat" : "-79.0000" })
    map.addMarker({"lon" : "45.5200", "lat" : "-122.6819" })
    map.addMarker({"lon" : "53.3478", "lat" : "-6.2597" })
    map.addMarker({"lon" : "1.3000", "lat" : "103.8000" })
   }
   else {
    map.clearMarkers()
   }
   marker =! marker
   screen.render()
}, 1000)

//set line charts dummy data

var transactionsData = {
   x: ['00:00', '00:05', '00:10', '00:15', '00:20', '00:30', '00:40', '00:50', '01:00', '01:10', '01:20', '01:30', '01:40', '01:50', '02:00', '02:10', '02:20', '02:30', '02:40', '02:50', '03:00', '03:10', '03:20', '03:30', '03:40', '03:50', '04:00', '04:10', '04:20', '04:30'],
   y: [0, 10, 40, 45, 45, 50, 55, 70, 65, 58, 50, 55, 60, 65, 70, 80, 70, 50, 40, 50, 60, 70, 82, 88, 89, 89, 89, 80, 72, 70]
}

var errorsData = {
   x: ['00:00', '00:05', '00:10', '00:15', '00:20', '00:25'],
   y: [30, 50, 70, 40, 50, 20]
}

setLineData(transactionsData, transactionsLine)
setLineData(errorsData, errorsLine)

function setLineData(mockData, line) {
  var last = mockData.y[mockData.y.length-1]
  mockData.y.shift()
  var num = Math.max(last + Math.round(Math.random()*10) - 5, 10)
  mockData.y.push(num)
  line.setData(mockData.x, mockData.y)
}

screen.key(['escape', 'q', 'C-c'], function(ch, key) {
  return process.exit(0);
});

screen.render()

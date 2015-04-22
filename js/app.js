'use strict';

/* App Module */

var serenoApp = angular.module('serenoApp', [
  'serenoAnimations',
  'serenoServices'
]);

serenoApp.controller('nodeController', function() {
	this.nodes = nodes;
});

var nodes = [
{
	name: "localhost",
	ip: "127.0.0.1",
	up: true
},
{
	name: "Planta1_printer",
	ip: "172.16.105.10",
	up: false
}
];
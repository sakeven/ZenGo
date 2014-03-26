ZenGo
=====
ZenGo is a Go package which function like a compiler of Zen Code but not at all.
##About Zen Coding
Emmet (previously known as **Zen Coding**) is a web-developer’s toolkit that can greatly improve your HTML & CSS workflow.

You can find more information about emmet on [Emmet Documentation](http://docs.emmet.io/) and [Github emmet](https://github.com/emmetio/emmet)
##Support
####E!
*!*if the element is a leaf element(like a leaf node of a tree(one of data structure))
#####html!
	<html></html>
####E#name!
#####div#name!
	<div id="name"></div>
####E.name!
#####div.name!
	<div class="name"></div>
#####div.one.two
	<div class="one two"></div>
#####div#name.one.two
	<div id="name" class="one two"></div>
####E>E!
#####head>p!
	<head>
	    <p></p>
	</head>
#####table>tr>td!
	<table>
		<tr>
    		<td></td>
		</tr>
	</table>
#####ul#name>li.item!
	<ul id="name">
    	<li class="item"></li>
	</ul>
####E!+E!
#####p!+p!
	<p></p>
	<p></p>
#####div#name>p.one!+p.two!
	<div id="name">
    	<p class="one"></p>
    	<p class="two"></p>
	</div>
####E>E!^+E!
*^* is a flag that we should return to the father element
#####div>h!^+div!
	<div>
		<h></h>
	</div>
	<div></div>
####E!*N
#####p!*3
	<p></p>
	<p></p>
	<p></p>
#####ul#name>li.item*3>a!
	<ul id="name">
    	<li class="item">
    		<a></a>
    	</li>
    	<li class="item">
    		<a></a>
    	</li>
    	<li class="item">
    		<a></a>
    	</li>
	</ul>
##Usage
##License
The package is under GPL2.0 license found int the [LICENSE](https://github.com/sakeven/ZenGo/blob/master/LICENSE) file.
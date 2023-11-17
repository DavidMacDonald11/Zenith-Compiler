#!/bin/bash

antlr='java -jar /home/david/bin/antlr-4.13.1-complete.jar'

$antlr -Dlanguage=Go ZenithLexer.g4 -o parser
$antlr -Dlanguage=Go ZenithParser.g4 -visitor -o parser

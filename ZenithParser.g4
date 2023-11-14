parser grammar ZenithParser;
options { tokenVocab = ZenithLexer; }

fileStat : expr EOF ;

expr : expr op=(TIMES | DIVIDE) NL? expr
     | expr op=(PLUS | MINUS) NL? expr
     | NUM
     ;

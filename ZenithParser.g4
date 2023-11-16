parser grammar ZenithParser;
options { tokenVocab = ZenithLexer; }

fileStat : expr EOF ;

expr : Left=expr Op=(TIMES | DIVIDE | REM) NL? Right=expr #mulExpr
     | Left=expr Op=(PLUS | MINUS) NL? Right=expr #addExpr
     | Num=NUM #numExpr
     ;

parser grammar ZenithParser;
options { tokenVocab = ZenithLexer; }

fileStat : expr EOF ;

expr : LPAREN NL? Inner=expr NL? RPAREN #parenExpr
     | Type=TYPE LPAREN NL? Inner=expr NL? RPAREN #castExpr
     | Op=(PLUS | MINUS) Right=expr #prefixExpr
     | Left=expr Op=(TIMES | DIVIDE | REM) NL? Right=expr #mulExpr
     | Left=expr Op=(PLUS | MINUS) NL? Right=expr #addExpr
     | Left=expr IF NL? Condition=expr NL? ELSE NL? Right=expr #ifExpr
     | Num=NUM #numExpr
     | Key=(TRUE | FALSE) #keyExpr
     ;

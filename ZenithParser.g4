parser grammar ZenithParser;
options { tokenVocab = ZenithLexer; }

fileStat : endedStat+
         | EOF
         ;

endedStat : NL? stat lineEnd ;

lineEnd : SEMICOLON
        | NL
        | EOF
        ;

stat : ID DEFINE_EQ expr #defineStat
     | expr #exprStat
     ;

expr : NUM #numExpr
     | ID #idExpr
     | Key=(TRUE | FALSE) #keyExpr
     | LPAREN NL? expr NL? RPAREN #parenExpr
     | TYPE LPAREN NL? expr NL? RPAREN #castExpr
     | Op=(PLUS | MINUS) Right=expr #prefixExpr
     | Left=expr Op=(TIMES | DIVIDE | REM) NL? Right=expr #mulExpr
     | Left=expr Op=(PLUS | MINUS) NL? Right=expr #addExpr
     | Left=expr IF NL? Condition=expr NL? ELSE NL? Right=expr #ifExpr
     ;

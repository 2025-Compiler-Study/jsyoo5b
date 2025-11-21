lexer grammar MiniCLexer;

// keywords

CONST   : 'const';
INT     : 'int';
VOID    : 'void';
IF      : 'if';
ELSE    : 'else';
WHILE   : 'while';
RETURN  : 'return';

// identifiers

IDENT   : ALPHABET
        | ('_' | ALPHABET) ('_' | ALPHABET | DEC_DIGIT)*;

// arithmetic operators

PLUS    : '+';
MINUS   : '-';
MUL     : '*';
DIV     : '/';
MOD     : '%';
INC     : '++';
DEC     : '--';

// assign operators

ASSIGN          : '=';
PLUS_ASSIGN     : '+=';
MINUS_ASSIGN    : '-=';
MUL_ASSIGN      : '*=';
DIV_ASSIGN      : '/=';
MOD_ASSIGN      : '%=';

// logical operators

LOGICAL_NOT : '!';
LOGICAL_AND : '&&';
LOGIC_OR    : '||';

// compare operators

EQUALS            : '==';
NOT_EQUALS        : '!=';
LESS              : '<';
LESS_OR_EQUALS    : '<=';
GREATER           : '>';
GREATER_OR_EQUALS : '>=';

// punctuation

L_PAREN     : '(';
R_PAREN     : ')';
L_CURLY     : '{';
R_CURLY     : '}';
L_BRACKET   : '[';
R_BRACKET   : ']';
COMMA       : ',';
SEMI        : ';';

// literal

DECIMAL_LIT : ('0' | [1-9] DEC_DIGIT*);
OCTAL_LIT   : '0' OCT_DIGIT+;
HEX_LIT     : '0' [xX] HEX_DIGIT+;

fragment DEC_DIGIT  : [0-9];
fragment OCT_DIGIT  : [0-7];
fragment HEX_DIGIT  : [0-9a-fA-F];
fragment ALPHABET   : [a-zA-Z];

// comments
WS          : [ \t\r\n]+    -> skip;
C89_COMMENT : '/*' .*? '*/' -> skip;
CPP_COMMENT : '//' ~[\r\n]* -> skip;

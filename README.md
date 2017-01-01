# A simple calculator

## Grammar
```
expr: term((ADD|MINUS)term)*
term: factor((MUL|DIV)factor)*
factor: number|number^number
number: INTEGER|LP expr RP|FUNC LP expr RP
```
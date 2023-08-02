# duskterpreter
 Interpreter made from Go

 ## How a "package"_test.go file works
 1. Start the lexer, which converts the string input into tokens with specific types
 2. Start the parser using the lexer output, which converts the tokens into usable datatypes (let statements, integer literals, return statements etc)
 3. Create the program using the parser as input
 4. Check for any parser errors using checkParserErrors(t, p)
 5. After this, all the other error catchers are used to make sure that the parsed content is correct to the answer key that we make. First  we create ok block statements which check that the ast types are correct starting from broad to specific (ExpressionStatement to PrefixExpression etc)
 6. Then we check if the actual values are correct, referencing off the answer key that we made. For example in a infix expression, we compare the left/right given values to the expected value and the opperator string to the expected string.

 Note: All the actual grunt work happens in lexer.go, parser.go and ast.go as they create the basis of the language using recusion, types and priorities.
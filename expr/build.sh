echo "//  DO NOT EDIT" > all.bnf
echo "//  This is generated file, see build.sh" >> all.bnf
echo "//  Sources: license.bnf lexer.bnf, expr.bnf, keyword.bnf" >> all.bnf
echo "" >> all.bnf
cat license.bnf >> all.bnf
cat lexer.bnf >> all.bnf
cat import.bnf >> all.bnf
cat expr.bnf >> all.bnf
cat keyword.bnf >> all.bnf
gocc all.bnf
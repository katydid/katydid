echo "//  DO NOT EDIT" > all.bnf
echo "//  This is generated file, see build.sh" >> all.bnf
echo "//  Sources: ../expr/license.bnf ../expr/lexer.bnf, lang.bnf, ../expr/expr.bnf, ../expr/keyword.bnf" >> all.bnf
echo "" >> all.bnf
cat ../expr/license.bnf >> all.bnf
cat ../expr/lexer.bnf >> all.bnf
cat lang.bnf >> all.bnf
cat ../expr/expr.bnf >> all.bnf
cat ../expr/keyword.bnf >> all.bnf
gocc all.bnf
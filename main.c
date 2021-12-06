#include <stdio.h>
#include <string.h>
#include <ctype.h>

#define RED "\033[0;31m"
#define YEL "\033[1;33m"
#define BRW "\033[0;33m"
#define NC "\033[0m"

#define ERR "\033[0;31m[ERROR]\033[0m"
#define WARN "\033[0;33m[WARN]\033[0m"

// Token Enum
enum t_token
{
    NONE,
    DONE,
    NUMBER,
    IDENTIFIER,
    KEYWORD,
    DELIMITER,
    OPERATOR,
    CHAR,
    STRING
};

// Constants
#define BUFFER_SIZE 256
#define SYM_TABLE_SIZE 1000

// Language Definitions
const char keywords[15][9] = {"array", "boolean", "char", "else", "false", "for", "function", "if",
                              "integer", "print", "return", "string", "true", "void", "while"};
const char whitespaces[5] = " \t\r\n";
const char delimiters[12] = "()[]{},;:";

// Operators which have two char form
#define INC_OP 900
#define DEC_OP 901
#define EQ_OP 902
#define GE_OP 903
#define LE_OP 904
#define NE_OP 905
#define AND_OP 906
#define OR_OP 907

// Symbol Lookup Table
char sym_table[SYM_TABLE_SIZE][BUFFER_SIZE];
int sym_count = 0;

int symbolLookup(char str[])
{
    for (int i = 0; i < sym_count; i++)
    {
        if (strcmp(str, sym_table[i]) == 0)
            return i;
    }
    return -1; //if not found return -1
}

int symbolInsert(char str[])
{
    strcpy(sym_table[sym_count++], str);
    return sym_count - 1;
}

// Strings Buffer
char str_buff[10000];
int lp_str_buff = 0; //last position in str_buff

// Testing Functions
int testKeyword(char str[])
{
    for (int i = 0; i < 15; i++)
    {
        if (strcmp(str, keywords[i]) == 0)
            return i + 1;
    }
    return 0; //if none matched return zero
}

int testGeneral(char c, char const charset[])
{
    for (int i = 0; charset[i]; i++)
    {
        if (c == charset[i])
            return 1;
    }
    return 0;
}

FILE *fp;
int line_no = 1;

void lex(int *tokenValue, enum t_token *tokenType)
{
    while (1)
    {
        //read char
        char ch = fgetc(fp);

        //handle eof
        if (ch == EOF)
        {
            *tokenValue = 0;
            *tokenType = DONE;
            return;
        }

        //handle newline
        if (ch == '\n')
        {
            line_no++;
            continue;
        }

        // handle whitespace
        if (testGeneral(ch, whitespaces))
        {
            continue;
        }

        // numbers
        if (isdigit(ch))
        {
            ungetc(ch, fp);
            fscanf(fp, "%d", tokenValue);
            *tokenType = NUMBER;
            return;
        }

        //identifier / keyword
        if (isalpha(ch) || ch == '_')
        {
            char buff[BUFFER_SIZE];
            int bi = 0;

            while (isalnum(ch) || ch == '_')
            {
                if (bi >= BUFFER_SIZE)
                {
                    printf("%s in line %d : identifiers cannot exceed length of %d\n", ERR, line_no, BUFFER_SIZE);
                }
                buff[bi++] = ch;
                ch = fgetc(fp);
            }
            buff[bi] = '\0'; //mark end of string
            ungetc(ch, fp);

            //keyword
            int k = testKeyword(buff);
            if (k)
            {
                *tokenValue = k;
                *tokenType = KEYWORD;
                return;
            }

            //identifier
            k = symbolLookup(buff);
            if (k == -1)
            {
                k = symbolInsert(buff);
            }
            *tokenValue = k;
            *tokenType = IDENTIFIER;
            return;
        }

        //char
        if (ch == '\'')
        {
            //read one character
            ch = fgetc(fp);
            //handle escaped characters
            if (ch == '\\')
            {
                char ch2 = fgetc(fp);
                if (ch2 == 'n')
                {
                    *tokenValue = 10;
                    *tokenType = CHAR;
                }
                else if (ch2 == '0')
                {
                    *tokenValue = 0;
                    *tokenType = CHAR;
                }
                else
                {
                    *tokenValue = ch2;
                    *tokenType = CHAR;
                    printf("%s in line %d : unsupported escaped character (\\%c)\n", WARN, line_no, ch2);
                }
            }
            else
            {
                *tokenValue = ch;
                *tokenType = CHAR;
            }
            //check ending
            ch = fgetc(fp);
            if (ch != '\'')
            {
                printf("%s in line %d : missing character closing.\n", ERR, line_no);
            }
            return;
        }

        //string
        if (ch == '\"')
        {
            *tokenValue = lp_str_buff;
            *tokenType = STRING;
            while (1)
            {
                char str_c = fgetc(fp);

                if (str_c == '\"')
                {
                    break;
                }

                if (str_c == EOF)
                {
                    printf("%s in line %d : missing string closing.\n", ERR, line_no);
                    *tokenType = DONE;
                    return;
                }

                if (str_c == '\n')
                {
                    line_no++;
                    printf("%s in line %d : missing string closing.\n", ERR, line_no);
                    break;
                }

                //handle escaped characters
                /*if(str_c == '\\'){
                    char ch2 = fgetc(fp);
                    if(ch2 == 'n'){
                        str_buff[lp_str_buff++] = '\n';
                    }else{
                        *tokenValue = ch2;
                        *tokenType = CHAR;
                        perror("warning in line %d : unsupported escaped character (\\%c)",line_no,ch2);
                    }
                }*/

                //append to strings buffer
                str_buff[lp_str_buff++] = str_c;
            }
            //when loop broke just mark string ending
            str_buff[lp_str_buff++] = '\0';
            //return token
            return;
        }

        //operators
        if (ch == '+')
        {
            char ch2 = fgetc(fp);
            if (ch2 == '+')
            {
                *tokenValue = INC_OP;
                *tokenType = OPERATOR;
                return;
            }
            else
                ungetc(ch2, fp);
            *tokenValue = '+';
            *tokenType = OPERATOR;
            return;
        }

        if (ch == '-')
        {
            char ch2 = fgetc(fp);
            if (ch2 == '-')
            {
                *tokenValue = DEC_OP;
                *tokenType = OPERATOR;
                return;
            }
            else
                ungetc(ch2, fp);
            *tokenValue = '-';
            *tokenType = OPERATOR;
            return;
        }

        /*if(ch == '='){
            char ch2 = fgetc(fp);
            if(ch2 == '=') {
                *tokenValue = EQ_OP;
                *tokenType = OPERATOR;
                return;
            }else ungetc(ch,fp);
            *tokenValue = '=';
            *tokenType = OPERATOR;
            return;
        }*/

        if (ch == '&')
        {
            char ch2 = fgetc(fp);
            if (ch2 == ch)
            {
                *tokenValue = AND_OP;
                *tokenType = OPERATOR;
                return;
            }
            else
            {
                printf("%s in line %d : syntax error.\n", ERR, line_no);
                continue;
            }
        }

        if (ch == '|')
        {
            char ch2 = fgetc(fp);
            if (ch2 == ch)
            {
                *tokenValue = OR_OP;
                *tokenType = OPERATOR;
                return;
            }
            else
            {
                printf("%s in line %d : syntax error.\n", ERR, line_no);
                continue;
            }
        }

        if (ch == '<' || ch == '>' || ch == '!' || ch == '=')
        {
            char ch2 = fgetc(fp);
            if (ch2 == '=')
            {
                switch (ch)
                {
                case '<':
                    *tokenValue = LE_OP;
                    break;
                case '>':
                    *tokenValue = GE_OP;
                    break;
                case '!':
                    *tokenValue = NE_OP;
                    break;
                case '=':
                    *tokenValue = EQ_OP;
                    break;
                }
                *tokenType = OPERATOR;
                return;
            }
            else
                ungetc(ch2, fp);

            *tokenValue = ch;
            *tokenType = OPERATOR;
            return;
        }

        if (ch == '/')
        {
            char ch2 = fgetc(fp);
            if (ch2 == '*')
            {
                // multi line comment
                while (1)
                {
                    ch = fgetc(fp);
                    if (ch == EOF)
                    {
                        printf("%s in line %d : multi line comment not closed.\n", ERR, line_no);
                        *tokenType = DONE;
                        return;
                    }
                    if (ch == '\n')
                        line_no++;
                    if (ch == '*')
                    {
                        char ch3 = fgetc(fp);
                        if (ch3 == '/')
                            break;
                    }
                }
                continue;
            }
            else if (ch2 == '/')
            {
                // single line comment
                while (1)
                {
                    ch = fgetc(fp);
                    if (ch == '\n')
                    {
                        line_no++;
                        break;
                    }
                    if (ch == EOF)
                    {
                        *tokenType = DONE;
                        return;
                    }
                }
                continue;
            }
            else
            {
                ungetc(ch2, fp);
                //division operator
                *tokenValue = ch;
                *tokenType = OPERATOR;
                return;
            }
        }

        if (ch == '*' || ch == '^' || ch == '%')
        {
            *tokenValue = ch;
            *tokenType = OPERATOR;
            return;
        }

        //delimiters
        if (testGeneral(ch, delimiters))
        {
            *tokenValue = ch;
            *tokenType = DELIMITER;
            return;
        }

        //if code reaches this part it means no matched rule which means invalid char
        printf("%s in line %d : invalid character (%c)\n", ERR, line_no, ch);
    }
}

void printToken(int *tokenValue, enum t_token *tokenType)
{
    switch (*tokenType)
    {
    case NUMBER:
        printf("NUM        %d\n", *tokenValue);
        break;
    case IDENTIFIER:
        printf("ID         %d\n", *tokenValue + 1);
        break;
    case KEYWORD:
        printf("KEYWORD    %d\n", *tokenValue);
        break;
    case CHAR:
        printf("CHAR       %d\n", *tokenValue);
        break;
    case STRING:
        printf("STRING     %d\n", *tokenValue);
        break;
    case OPERATOR:
    case DELIMITER:
        if (*tokenValue < 900)
        {
            printf("%c          NONE\n", (char)*tokenValue);
        }
        else if (*tokenValue == INC_OP)
            printf("++         NONE\n");
        else if (*tokenValue == DEC_OP)
            printf("--         NONE\n");
        else if (*tokenValue == EQ_OP)
            printf("==         NONE\n");
        else if (*tokenValue == GE_OP)
            printf(">=         NONE\n");
        else if (*tokenValue == LE_OP)
            printf("<=         NONE\n");
        else if (*tokenValue == NE_OP)
            printf("!=         NONE\n");
        else if (*tokenValue == AND_OP)
            printf("&&         NONE\n");
        else if (*tokenValue == OR_OP)
            printf("||         NONE\n");
        break;
    }
}

int main(int argc, char *argv[])
{
    fp = fopen(argv[1], "r");
    if (!fp)
    {
        printf("%s failed to read file!", ERR);
        return -1;
    }
    else
    {
        printf("%s started compiling %s ... %s\n", YEL, argv[1], NC);
    }

    printf("%s performing lexical analysis ... %s\n", YEL, NC);
    enum t_token tokenType = NONE;
    int tokenValue;
    while (tokenType != DONE)
    {
        lex(&tokenValue, &tokenType);
        printToken(&tokenValue, &tokenType);
    }

    fclose(fp);
    return 0;
}

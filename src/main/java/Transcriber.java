public class Transcriber{
    private static String inputString;

    public Transcriber(String input){
        this.inputString = input;
    }

    public String getTranscription(){
        return transcribe();
    }
    private String transcribe(){
        String str = this.inputString.toLowerCase();
        char[] inpArr = str.toCharArray();
        StringBuilder builder = new StringBuilder();
        builder.append("\uD800\uDFD0");

        for(int i = 0; i < inpArr.length; i++) {
            char ch = inpArr[i];
            char next;
            if (i < inpArr.length - 1) {
                next = Character.toLowerCase(inpArr[i + 1]);
            } else {
                next = '\0';
            }
            if (ch == 'ā' || ch == 'a' && (i == 0 || next == '\'')) builder.append("\uD800\uDFA0");
            if (ch == 'b') builder.append("\uD800\uDFB2");
            if (ch == 'ç' || ch == 'c' && next == '\'') {builder.append("\uD800\uDFC2"); continue;}
            if (ch == 'c') builder.append("\uD800\uDFA8");
            if (ch == 'd') {
                if (next == 'i') builder.append("\uD800\uDFAE");
                else if (next == 'u') builder.append("\uD800\uDFAF");
                else builder.append("\uD800\uDFAD");
            }
            if (ch == 'f') builder.append("\uD800\uDFB3");
            if (ch == 'g') {
                if (next == 'u') builder.append("\uD800\uDFA6");
                else builder.append("\uD800\uDFA5");
            }
            if (ch == 'h') builder.append("\uD800\uDFC3");
            if (ch == 'i') builder.append("\uD800\uDFA1");
            if (ch == 'j') {
                if (next == 'i') builder.append("\uD800\uDFA6");
                else builder.append("\uD800\uDFA9");
            }
            if (ch == 'k') {
                if (next == 'u') builder.append("\uD800\uDFA4");
                else builder.append("\uD800\uDFA3");
            }
            if (ch == 'l') builder.append("\uD800\uDFBE");
            if (ch == 'm') {
                if (next == 'i') builder.append("\uD800\uDFB7");
                else if (next == 'u') builder.append("\uD800\uDFB8");
                else builder.append("\uD800\uDFB6");
            }
            if (ch == 'n') {
                if (next == 'u') builder.append("\uD800\uDFB5");
                else builder.append("\uD800\uDFB4");
            }
            if (ch == 'p') builder.append("\uD800\uDFB1");
            if (ch == 'r') {
                if (next == 'u') builder.append("\uD800\uDFBD");
                else builder.append("\uD800\uDFBC");
            }
            if (ch == 'š' || ch == 's' && next == '\'') {builder.append("\uD800\uDFC1"); continue;}
            if (ch == 's') builder.append("\uD800\uDFBF");
            if (ch == 'θ' || ch == 't' && next == '\'') {builder.append("\uD800\uDFB0"); continue;}
            if (ch == 't') {
                if (next == 'u') builder.append("\uD800\uDFAC");
                else builder.append("\uD800\uDFAB");
            }
            if (ch == 'u') builder.append("\uD800\uDFA2");
            if (ch == 'v') {
                if (next == 'i') builder.append("\uD800\uDFBB");
                else builder.append("\uD800\uDFBA");
            }
            if (ch == 'x') builder.append("\uD800\uDFA7");
            if (ch == 'y') builder.append("\uD800\uDFB9");
            if (ch == 'z') builder.append("\uD800\uDFC0");
            if (ch == ' ') {
                if (next == '\n') {
                    builder.append(" \n\uD800\uDFD0");
                    i++;
                } else {
                    builder.append(" \uD800\uDFD0");
                }
            }
            if (ch == '\n') builder.append("\n");
            if (ch == '1') {
                if (i == inpArr.length - 1 || next == ' ') builder.append("\uD800\uDFD1");
                else if (i + 1 == inpArr.length - 1 || inpArr[i + 2] == ' ') builder.append("\uD800\uDFD3");
                else builder.append("\uD800\uDFD5");
            }
            if (ch == '2') {
                if (i == inpArr.length - 1 || next == ' ') builder.append("\uD800\uDFD2");
                else builder.append("\uD800\uDFD4");
            }
            if (ch == '3') {
                if (i == inpArr.length - 1 || next == ' ') builder.append("\uD800\uDFD2\uD800\uDFD1");
                else builder.append("\uD800\uDFD4\uD800\uDFD3");
            }
            if (ch == '4') {
                if (i == inpArr.length - 1 || next == ' ') builder.append("\uD800\uDFD2\uD800\uDFD2");
                else builder.append("\uD800\uDFD4\uD800\uDFD4");
            }
            if (ch == '5') {
                if (i == inpArr.length - 1 || next == ' ') builder.append("\uD800\uDFD2\uD800\uDFD2\uD800\uDFD1");
                else builder.append("\uD800\uDFD4\uD800\uDFD4\uD800\uDFD3");
            }
            if (ch == '6') {
                if (i == inpArr.length - 1 || next == ' ') builder.append("\uD800\uDFD2\uD800\uDFD2\uD800\uDFD2");
                else builder.append("\uD800\uDFD4\uD800\uDFD4\uD800\uDFD4");
            }
            if (ch == '7') {
                if (i == inpArr.length - 1 || next == ' ') builder.append("\uD800\uDFD2\uD800\uDFD2\uD800\uDFD2\uD800\uDFD1");
                else builder.append("\uD800\uDFD4\uD800\uDFD4\uD800\uDFD4\uD800\uDFD3");
            }
            if (ch == '8') {
                if (i == inpArr.length - 1 || next == ' ') builder.append("\uD800\uDFD2\uD800\uDFD2\uD800\uDFD2\uD800\uDFD2");
                else builder.append("\uD800\uDFD4\uD800\uDFD4\uD800\uDFD4\uD800\uDFD4");
            }
            if (ch == '9') {
                if (i == inpArr.length - 1 || next == ' ') builder.append("\uD800\uDFD2\uD800\uDFD2\uD800\uDFD2\uD800\uDFD2\uD800\uDFD1");
                else builder.append("\uD800\uDFD4\uD800\uDFD4\uD800\uDFD4\uD800\uDFD4\uD800\uDFD3");
            }
        }

        builder.append("\n");
        if (str.contains("baga")) builder.append("\nLogogram for baga: \uD800\uDFCE");
        if (str.contains("ahuramazda'") || str.contains("ahuramazdā")) builder.append("\nLogograms for Ahuramazdā: \uD800\uDFC8 / \uD800\uDFC9 / \uD800\uDFCA");
        if (str.contains("dahya'u") || str.contains("dahyāu")) builder.append("\nLogograms for dahyāu: \uD800\uDFCC / \uD800\uDFCD");
        if (str.contains("xs'a'yat'iya") || str.contains("xšāyaθiya")) builder.append("\nLogogram for xšāyaθiya: \uD800\uDFCB");
        if (str.contains("θātiy dārayavauš xšāyaθiya vašnā ahuramazdāha adam xšāyaθiya amiy ahuramazdāha xšaçam manā frābara"))
            builder.append("\n\nBy the way, it means \"King Darius says: By the grace of Ahuramazda am I king; Ahuramazda has granted me the kingdom.\"");

        return builder.toString();
    }
}
import org.glassfish.jersey.jaxb.internal.XmlCollectionJaxbProvider;
import org.telegram.telegrambots.bots.TelegramLongPollingBot;
import org.telegram.telegrambots.meta.TelegramBotsApi;
import org.telegram.telegrambots.meta.api.methods.send.SendMessage;
import org.telegram.telegrambots.meta.api.objects.Message;
import org.telegram.telegrambots.meta.api.objects.Update;
import org.telegram.telegrambots.meta.api.objects.replykeyboard.ReplyKeyboardMarkup;
import org.telegram.telegrambots.meta.api.objects.replykeyboard.buttons.KeyboardButton;
import org.telegram.telegrambots.meta.api.objects.replykeyboard.buttons.KeyboardRow;
import org.telegram.telegrambots.meta.exceptions.TelegramApiException;
import org.telegram.telegrambots.updatesreceivers.DefaultBotSession;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Properties;

public class Main extends TelegramLongPollingBot {
    public static void main(String[] args) {
        try {
            TelegramBotsApi telegramBotsApi = new TelegramBotsApi(DefaultBotSession.class);
            telegramBotsApi.registerBot(new Main());
        } catch (TelegramApiException e) {
            e.printStackTrace();
        }
    }

    private void sendMsg(Message message, String text) {
        SendMessage sendMessage = new SendMessage();
        sendMessage.enableMarkdown(true);
        sendMessage.setChatId(message.getChatId().toString());
        sendMessage.setReplyToMessageId(message.getMessageId());
        sendMessage.setText(text);
        try{
            setButtons(sendMessage);
            execute(sendMessage);
        } catch (TelegramApiException e) {
            e.printStackTrace();
        }
    }

    @Override
    public void onUpdateReceived(Update update) {
        Message message = update.getMessage();
        if(message != null && message.hasText()){
            String result = answer(message.getText());
            sendMsg(message, result);
        }
    }

    @Override
    public String getBotUsername() {
        return "OldPersianTranscriberBot";
    }

    @Override
    public String getBotToken() {
        String token = "";
        Properties prop = new Properties();
        try {
            //load a properties file from class path, inside static method
            prop.load(XmlCollectionJaxbProvider.App.class.getClassLoader().getResourceAsStream("config.properties"));

            //get the property value and print it out
            token = prop.getProperty("token");
        } catch(IOException ex) {
            ex.printStackTrace();
        }
        return token;
    }

    private void setButtons(SendMessage sendMessage) {
        ReplyKeyboardMarkup replyKeyboardMarkup = new ReplyKeyboardMarkup();
        sendMessage.setReplyMarkup(replyKeyboardMarkup);
        replyKeyboardMarkup.setSelective(true);
        replyKeyboardMarkup.setResizeKeyboard(true);
        replyKeyboardMarkup.setOneTimeKeyboard(false);

        List<KeyboardRow> keyboardRowList = new ArrayList<>();
        KeyboardRow keyboardRow1 = new KeyboardRow();

        keyboardRow1.add(new KeyboardButton("/start"));
        keyboardRow1.add(new KeyboardButton("/example"));
        keyboardRow1.add(new KeyboardButton("/help"));

        keyboardRowList.add(keyboardRow1);
        replyKeyboardMarkup.setKeyboard(keyboardRowList);
    }

    private String answer(String input) {
        switch (input) {
            case("/start"): return
                    "Hello. Let's transcribe your phrase into the Old Persian cuneiform." +
                            "\nIf you just want to try it, click /example, copy the® phrase from the Behistun Inscription and send it back to the bot." +
                            "\nTo learn some more info click /help." +
                            "\nMade by @epaolinos (Pavel Egizaryan) in 2019.";
            case("/help"): return
                    "You can use digraphs instead of special characters:\n" +
                            "\nā = a' " +
                            "\nç = c' " +
                            "\nθ = t' " +
                            "\nš = s'" +
                            "\n\nThe numbers for 200 and more are not known." +
                            "\nMake sure your words and numbers are correctly separated with spaces." +
                            "\nFor some words you will see possible logograms under the normal transcription.";
            case("/example"): return
                    "\n\nθātiy Dārayavauš xšāyaθiya vašnā Ahuramazdāha adam xšāyaθiya amiy Ahuramazdāha xšaçam manā frābara";
        }

        //return "Ok";

        Transcriber transcriber = new Transcriber(input);
        return transcriber.getTranscription();
    }
}

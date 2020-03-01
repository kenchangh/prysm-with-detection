from telegram import Bot
from telegram.ext import Updater, CommandHandler
from flask import Flask, request
app = Flask(__name__)

bot_token = '1110238566:AAGWRj_8MQVLohYvDYSgbRBcSv3TD8BFnQA'
bot = Bot(token=bot_token)
updater = Updater(token=bot_token, use_context=True)

def start(update, context):
    welcome = "Welcome to ETH2.0 Staking Monitor. To start, enter the pubkey you want to monitor."
    context.bot.send_message(chat_id=update.effective_chat.id, text=welcome)

def input_pubkey(update, context):
    msg = "Thank you, we are now tracking your validator's activity."
    context.bot.send_message(chat_id=update.effective_chat.id, text=msg)
    print(update.effective_chat.id)

start_handler = CommandHandler('start', start)
track_handler = CommandHandler('track', input_pubkey)

dispatcher = updater.dispatcher
updater.dispatcher.add_handler(start_handler)
updater.dispatcher.add_handler(track_handler)
updater.start_polling()

@app.route('/alert/<mode>', methods=['POST'])
def alert_user(mode):
    slot = request.data

    if mode == 'telegram':
        # bot.send_message(chat_id=chat_id, text="I'm sorry Dave I'm afraid I can't do that.")
        bot.send_message(chat_id=324513231, text=slot.decode('utf-8'))

    return 'Hello, World!'

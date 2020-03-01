from telegram import Bot
from telegram.ext import Updater, CommandHandler
from flask import Flask, request
app = Flask(__name__)

bot_token = '1110238566:AAGWRj_8MQVLohYvDYSgbRBcSv3TD8BFnQA'
bot = Bot(token=bot_token)
updater = Updater(token=bot_token, use_context=True)

notify_keys = {}

def start(update, context):
    global notify_keys
    welcome = "Welcome to ETH2.0 Staking Monitor. To start, enter the pubkey you want to monitor."
    context.bot.send_message(chat_id=update.effective_chat.id, text=welcome)

def input_pubkey(update, context):
    global notify_keys
    msg = "Thank you, we are now tracking your validator's activity."
    context.bot.send_message(chat_id=update.effective_chat.id, text=msg)
    msg_text = update['message']['text']
    _, pubkey = msg_text.split(' ')

    subscribers = notify_keys.get(pubkey)
    if not subscribers:
        notify_keys[pubkey] = set([])
    
    notify_keys[pubkey].add(update.effective_chat.id)

start_handler = CommandHandler('start', start)
track_handler = CommandHandler('track', input_pubkey, pass_args=True)

dispatcher = updater.dispatcher
updater.dispatcher.add_handler(start_handler)
updater.dispatcher.add_handler(track_handler)
updater.start_polling()

@app.route('/alert/<mode>', methods=['POST'])
def alert_user(mode):
    global notify_keys
    data = request.json
    slot = data['slot']
    pubkey = data['pubkey']

    if mode == 'telegram':
        subscribers = notify_keys.get(pubkey)
        print(f'Received event for {pubkey} for subscribers {subscribers}')
        if not subscribers:
            return ''
        
        msg = f'Double vote happened at slot {slot} with validator {pubkey}'
        
        for subscriber in subscribers:
            bot.send_message(chat_id=subscriber, text=msg)

    return 'Hello, World!'

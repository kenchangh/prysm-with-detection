from flask import Flask, request
app = Flask(__name__)


@app.route('/alert/<mode>', methods=['POST'])
def alert_user(mode):
    slot = request.data

    if mode == 'telegram':
        print(mode, slot)

    return 'Hello, World!'

import os
import json
from flask import Flask, request, render_template, jsonify
import random

# height = 0

# with open('watermark.json') as json_file:
#     data = json.load(json_file)
#     height = data["height"]

app = Flask(__name__)

validators = {}

selected_attestations = {}


def dir_last_updated(folder):
    return str(max(os.path.getmtime(os.path.join(root_path, f))
                   for root_path, dirs, files in os.walk(folder)
                   for f in files))


@app.route('/')
def main():
    return render_template('index.html', last_updated=dir_last_updated('static'))


@app.route('/validator')
def load_val():
    global validators
    global selected_attestations
    return jsonify({
        "validators": validators,
        "selected": selected_attestations
    })


@app.route('/update', methods=['POST'])
def update_height():
    global validators
    global selected_attestations

    data = request.json
    validator_id = data['validatorID']
    dataStr = data['dataStr']
    slot = data['slot']

    validator_exists = validators.get(validator_id)
    if not validator_exists:
        validators[validator_id] = {}
    validators[validator_id][slot] = data

    if slot not in selected_attestations:
        selected_attestations[slot] = data
        return '1'
    else:
        if dataStr != selected_attestations[slot]['dataStr']:
            return '0'
        else:
            return '1'

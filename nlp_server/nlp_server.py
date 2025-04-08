import spacy
from flask import Flask, request, jsonify

nlp = spacy.load('en_core_web_sm')

app = Flask(__name__)

@app.route('/analyze', methods=['POST'])
def analyze_text():
    text = request.json.get('text')
    doc = nlp(text)
    result = {
        "tokens": [token.text for token in doc],
        "entities": [(ent.text, ent.label_) for ent in doc.ents],
        "pos_tags": [(token.text, token.pos_) for token in doc],
    }
    return jsonify(result)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)

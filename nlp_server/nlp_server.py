import spacy
from flask import Flask, request, jsonify

app = Flask(__name__)
nlp = spacy.load("en_core_web_sm")

INTENT_RULES = {
    "refund": [
        "refund", "return", "money back"
    ],
    "orderStatus": [
        "order status", "where is my order", "track", "delivery"
    ],
    "greeting": [
        "hi", "hello", "hey", "good morning", "good evening"
    ],
    "goodbye": [
        "bye", "goodbye", "see you", "thanks"
    ],
    "cancelOrder": [
        "cancel", "cancel my order", "stop order"
    ],
    "changeAddress": [
        "change address", "wrong address", "update shipping"
    ],
    "technicalIssue": [
        "error", "bug", "not working", "crash", "issue", "problem"
    ],
    "paymentIssue": [
        "payment", "card declined", "checkout failed", "can’t pay"
    ],
    "speakToHuman": [
        "talk to a person", "human", "agent", "representative"
    ],
    "productInfo": [
        "product info", "tell me about", "details", "specs"
    ],
    "businessHours": [
        "business hours", "working hours", "when are you open"
    ],
    "supportContact": [
        "contact support", "support team", "how to contact support"
    ]
}


def extract_lemmas(doc):
    return [token.lemma_.lower() for token in doc]

def detect_intent(text: str) -> str:
    doc = nlp(text.lower())
    lemmas = extract_lemmas(doc)
    lowered_text = text.lower()

    for intent, keywords in INTENT_RULES.items():
        if any(keyword in lemmas for keyword in keywords) or any(keyword in lowered_text for keyword in keywords):
            return intent

    return "unknown"

@app.route('/intent', methods=['POST'])
def get_intent():
    data = request.get_json()
    if not data or "text" not in data:
        return jsonify({"error": "Missing 'text' field"}), 400

    intent = detect_intent(data["text"])
    return jsonify({"intent": intent})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
from flask import Flask, render_template
import requests

app = Flask(__name__)

@app.route("/")
def home():
    try:
        response = requests.get("http://system_info:5000/info")
        info = response.json()
    except Exception as e:
        info = {"error": str(e)}
    return render_template("index.html", info=info)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5001)

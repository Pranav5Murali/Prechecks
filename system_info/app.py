from flask import Flask, jsonify
import platform
import psutil

app = Flask(__name__)

@app.route("/info", methods=["GET"])
def system_info():
    info = {
        "platform": platform.system(),
        "platform_version": platform.version(),
        "cpu_count": psutil.cpu_count(),
        "memory": psutil.virtual_memory().total // (1024 * 1024),
    }
    return jsonify(info)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)

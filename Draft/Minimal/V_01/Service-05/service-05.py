from flask import Flask, request, jsonify
import pandas as pd
from sklearn.linear_model import LinearRegression

app = Flask(__name__)

@app.route('/build-model', methods=['POST'])
def build_model():
    data = request.get_json()
    df = pd.DataFrame(data)


    return jsonify(df)

if __name__ == '__main__':
    app.run(port=8085)

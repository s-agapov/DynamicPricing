from flask import Flask, request, jsonify
import pandas as pd
from sklearn.linear_model import LinearRegression

app = Flask(__name__)

@app.route('/build-model', methods=['POST'])
def build_model():
    data = request.get_json()
    df = pd.DataFrame(data)

    X = df[['column1', 'column2']]
    y = df['target']

    model = LinearRegression()
    model.fit(X, y)

    coef = model.coef_.tolist()
    intercept = model.intercept_

    return jsonify({'coef': coef, 'intercept': intercept})

if __name__ == '__main__':
    app.run(port=8083)

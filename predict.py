from flask import Flask, request, jsonify
from PIL import Image
from img2vec_pytorch import Img2Vec
import torch
import torch.nn as nn

app = Flask(__name__)
img2vec = Img2Vec(cuda=torch.cuda.is_available())

device = torch.device('cuda' if torch.cuda.is_available() else 'cpu')

class MLP(nn.Module):
    def __init__(self, input_dim=512, output_dim=120):
        super().__init__()
        self.net = nn.Sequential(
            nn.Linear(input_dim, 2048),
            nn.ReLU(),
            nn.Linear(2048, 1024),
            nn.ReLU(),
            nn.Linear(1024, 512),
            nn.ReLU(),
            nn.Linear(512, 256),
            nn.Tanh(),
            nn.Linear(256, 128),
            nn.Tanh(),
            nn.Linear(128, 256),
            nn.Tanh(),
            nn.Linear(256, output_dim),
            nn.LogSoftmax(dim=1)
        )

    def forward(self, x):
        return self.net(x)

# Load the model
model = torch.load('ai/full_best_model.pth', map_location=device)
model.to(device)
model.eval()

@app.route('/predict', methods=['POST'])
def predict_api():
    file = request.files['getImg']
    image = Image.open(file.stream).convert('RGB')
    print("Received image:", file.filename)

    vec = img2vec.get_vec(image, tensor=True).flatten().unsqueeze(0).to(device)

    with torch.no_grad():
        output = model(vec)
        predicted_class = torch.argmax(output, dim=1).item()
        confidence = torch.softmax(output, dim=1)[0][predicted_class].item()

        print([predicted_class, confidence])

        return jsonify({
            "class": predicted_class,
            "confidence": round(confidence, 4)
        })

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)

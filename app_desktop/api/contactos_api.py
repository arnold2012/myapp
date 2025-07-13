import requests

BASE_URL = "http://localhost:8080"

def obtener_contactos():
    try:
        response = requests.get(f"{BASE_URL}/api/contactos")
        response.raise_for_status()
        return response.json()  # Asumiendo que devuelve un JSON válido
    except Exception as e:
        print("Error al obtener contactos:", e)
        return []

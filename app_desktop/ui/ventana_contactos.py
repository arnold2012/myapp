from PyQt6.QtWidgets import (
    QWidget, QVBoxLayout, QLabel, QTableWidget, QTableWidgetItem
)
from api.contactos_api import obtener_contactos

class VentanaContactos(QWidget):
    def __init__(self):
        super().__init__()
        self.setWindowTitle("Lista de Contactos")
        self.resize(600, 400)

        self.layout = QVBoxLayout(self)

        self.label = QLabel("Contactos disponibles:")
        self.layout.addWidget(self.label)

        self.tabla = QTableWidget()
        self.layout.addWidget(self.tabla)

        self.cargar_contactos()

def cargar_contactos(self):
    contactos = obtener_contactos()
    print("Contactos recibidos:", contactos)  # <-- Agrega esto

    self.tabla.setRowCount(len(contactos))
    self.tabla.setColumnCount(3)
    self.tabla.setHorizontalHeaderLabels(["Nombre", "Documento", "Teléfono"])

    for fila, contacto in enumerate(contactos):
        self.tabla.setItem(fila, 0, QTableWidgetItem(contacto.get("nombre_razon", "")))
        self.tabla.setItem(fila, 1, QTableWidgetItem(contacto.get("nro_documento", "")))
        self.tabla.setItem(fila, 2, QTableWidgetItem(contacto.get("telefono", "")))

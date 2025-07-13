import sys
from PyQt6.QtWidgets import QApplication
from ui.ventana_contactos import VentanaContactos

if __name__ == "__main__":
    app = QApplication(sys.argv)
    ventana = VentanaContactos()
    ventana.show()
    sys.exit(app.exec())

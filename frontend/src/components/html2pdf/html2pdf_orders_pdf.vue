<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div v-if="cargando" class="bg-white p-6 rounded-lg shadow-lg text-center">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500 mx-auto mb-4"></div>
      <p class="text-lg font-medium">Generando PDF...</p>
    </div>
    <div v-if="error" class="bg-white p-6 rounded-lg shadow-lg max-w-md">
      <h3 class="text-lg font-bold text-red-600 mb-2">Error al generar PDF</h3>
      <p class="mb-4">{{ error }}</p>
      <button 
        @click="emit('cerrar')" 
        class="px-4 py-2 bg-gray-300 text-gray-800 rounded hover:bg-gray-400"
      >
        Cerrar
      </button>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref, computed } from 'vue'
import ordersApi from '@/api/orders'

const props = defineProps({
  orderId: {
    type: Number,
    required: true
  }
})

const emit = defineEmits(['cerrar'])
const cargando = ref(true)
const error = ref(null)
const items = ref([])
const orderData = ref(null)

const totales = computed(() => {
  let totalGravado10 = 0
  let totalGravado5 = 0
  let totalExenta = 0
  let totalIVA = 0
  let totalGeneral = 0
  
  items.value.forEach(item => {
    totalGravado10 += item.gravado_10 || 0
    totalGravado5 += item.gravado_5 || 0
    totalExenta += item.exenta || 0
    totalIVA += item.iva_calculado || 0
    totalGeneral += item.total_a_pagar || 0
  })
  
  return {
    totalGravado10,
    totalGravado5,
    totalExenta,
    totalIVA,
    totalGeneral
  }
})

const generarPDF = async () => {
  console.log('Generando PDF para la orden ID:', props.orderId)
  
  try {
    if (!props.orderId) {
      throw new Error('No se proporcionó un ID de orden válido')
    }

    const res = await ordersApi.getOrderForPrint(props.orderId)
    console.log('Respuesta de getOrderForPrint:', res)
    
    if (!res.success) {
      throw new Error(res.error || 'Error al obtener datos de la orden')
    }

    const data = res.data
    if (!data || data.length === 0) {
      throw new Error('No hay datos para generar el PDF')
    }

    console.log('Datos de la orden obtenidos correctamente:', data)
    
    items.value = data
    orderData.value = data[0]
    
    // Esperar a que el DOM se actualice
    await new Promise(resolve => setTimeout(resolve, 100))
    
    // Create a simple HTML string instead of using DOM elements
    const htmlContent = `
      <html>
        <head>
          <style>
            body {
              font-family: 'Helvetica', 'Arial', sans-serif;
              margin: 0;
              padding: 0;
              color: #333;
              line-height: 1.5;
            }
            .container {
              max-width: 800px;
              margin: 0 auto;
              padding: 40px;
            }
            .header {
              text-align: center;
              margin-bottom: 30px;
              padding-bottom: 20px;
              border-bottom: 2px solid #4a5568;
            }
            .header h1 {
              font-size: 28px;
              color: #2d3748;
              margin: 0;
              padding: 0;
              text-transform: uppercase;
              letter-spacing: 1px;
            }
            .document-info {
              display: flex;
              justify-content: space-between;
              margin-bottom: 30px;
            }
            .document-details, .client-details {
              width: 48%;
            }
            .section-title {
              font-size: 16px;
              text-transform: uppercase;
              color: #4a5568;
              margin-bottom: 10px;
              font-weight: bold;
              border-bottom: 1px solid #e2e8f0;
              padding-bottom: 5px;
            }
            .info-row {
              margin-bottom: 5px;
              display: flex;
            }
            .info-label {
              font-weight: bold;
              width: 140px;
            }
            .info-value {
              flex: 1;
            }
            .client-box {
              background-color: #f8fafc;
              border: 1px solid #e2e8f0;
              border-radius: 5px;
              padding: 15px;
            }
            table {
              width: 100%;
              border-collapse: collapse;
              margin-bottom: 30px;
            }
            th {
              background-color: #f1f5f9;
              color: #4a5568;
              font-weight: bold;
              text-align: left;
              padding: 12px;
              border: 1px solid #e2e8f0;
            }
            td {
              padding: 10px 12px;
              border: 1px solid #e2e8f0;
              vertical-align: top;
            }
            .text-right {
              text-align: right;
            }
            .totals-table {
              width: auto;
              margin-left: auto;
              margin-bottom: 40px;
            }
            .totals-table td {
              padding: 8px 12px;
            }
            .totals-table .total-row {
              font-weight: bold;
              border-top: 2px solid #4a5568;
            }
            .totals-table .total-row td {
              padding-top: 10px;
            }
            .signatures {
              display: flex;
              justify-content: space-between;
              margin-top: 60px;
            }
            .signature-box {
              width: 40%;
              text-align: center;
            }
            .signature-line {
              border-top: 1px solid #4a5568;
              padding-top: 10px;
              font-size: 14px;
            }
            .footer {
              margin-top: 60px;
              text-align: center;
              font-size: 12px;
              color: #718096;
              border-top: 1px solid #e2e8f0;
              padding-top: 20px;
            }
            .item-row:nth-child(even) {
              background-color: #f8fafc;
            }
            .document-number {
              font-size: 18px;
              color: #2d3748;
              margin-top: 5px;
              font-weight: bold;
            }
            .logo-space {
              height: 60px;
              margin-bottom: 15px;
              display: flex;
              justify-content: center;
              align-items: center;
            }
            .logo-placeholder {
              font-style: italic;
              color: #a0aec0;
              border: 1px dashed #cbd5e0;
              padding: 10px 20px;
              display: inline-block;
            }
          </style>
        </head>
        <body>
          <div class="container">
            <div class="header">
              <div class="logo-space">
                <span class="logo-placeholder">Logo de la Empresa</span>
              </div>
              <h1>Orden de Venta</h1>
              <div class="document-number">N° ${orderData.value.numero_order}</div>
            </div>
            
            <div class="document-info">
              <div class="document-details">
                <div class="section-title">Detalles del Documento</div>
                <div class="info-row">
                  <div class="info-label">Fecha:</div>
                  <div class="info-value">${formatDate(orderData.value.fecha_expedicion)}</div>
                </div>
                <div class="info-row">
                  <div class="info-label">Condición de Pago:</div>
                  <div class="info-value">${orderData.value.describe_condicion}</div>
                </div>
              </div>
              
              <div class="client-details">
                <div class="section-title">Información del Cliente</div>
                <div class="client-box">
                  <div class="info-row">
                    <div class="info-label">Razón Social:</div>
                    <div class="info-value">${orderData.value.razon_social}</div>
                  </div>
                  <div class="info-row">
                    <div class="info-label">RUC/CI:</div>
                    <div class="info-value">${orderData.value.ruc_ci || 'N/A'}</div>
                  </div>
                </div>
              </div>
            </div>
            
            <div class="section-title">Detalle de Productos</div>
            <table>
              <thead>
                <tr>
                  <th style="width: 15%;">Código</th>
                  <th style="width: 40%;">Descripción</th>
                  <th style="width: 15%;" class="text-right">Cantidad</th>
                  <th style="width: 15%;" class="text-right">Precio</th>
                  <th style="width: 15%;" class="text-right">Total</th>
                </tr>
              </thead>
              <tbody>
                ${items.value.map((item, index) => `
                  <tr class="item-row">
                    <td>${item.cod_item}</td>
                    <td>${item.descripcion_item}</td>
                    <td class="text-right">${item.cantidad}</td>
                    <td class="text-right">${formatNumber(item.precio_unitario)}</td>
                    <td class="text-right">${formatNumber(item.total_item)}</td>
                  </tr>
                `).join('')}
              </tbody>
            </table>
            
            <table class="totals-table">
              <tr>
                <td class="text-right">Gravado 10%:</td>
                <td class="text-right">${formatNumber(totales.value.totalGravado10)}</td>
              </tr>
              <tr>
                <td class="text-right">Gravado 5%:</td>
                <td class="text-right">${formatNumber(totales.value.totalGravado5)}</td>
              </tr>
              <tr>
                <td class="text-right">Exenta:</td>
                <td class="text-right">${formatNumber(totales.value.totalExenta)}</td>
              </tr>
              <tr>
                <td class="text-right">IVA:</td>
                <td class="text-right">${formatNumber(totales.value.totalIVA)}</td>
              </tr>
              <tr class="total-row">
                <td class="text-right">TOTAL:</td>
                <td class="text-right">${formatNumber(totales.value.totalGeneral)}</td>
              </tr>
            </table>
            
            <div class="signatures">
              <div class="signature-box">
                <div class="signature-line">Firma del Cliente</div>
              </div>
              <div class="signature-box">
                <div class="signature-line">Firma Autorizada</div>
              </div>
            </div>
            
            <div class="footer">
              <p>Este documento es una constancia de la orden de venta.</p>
              <p>Gracias por su preferencia.</p>
            </div>
          </div>
        </body>
      </html>
    `;
    
    // Use jsPDF directly
    const { jsPDF } = await import('jspdf');
    const doc = new jsPDF();
    
    // Use html2canvas to render the HTML string
    const { default: html2canvas } = await import('html2canvas');
    
    // Create a temporary div to hold our HTML
    const tempDiv = document.createElement('div');
    tempDiv.innerHTML = htmlContent;
    document.body.appendChild(tempDiv);
    
    try {
      const canvas = await html2canvas(tempDiv, {
        scale: 2,
        useCORS: true,
        logging: false,
        backgroundColor: '#ffffff'
      });
      
      const imgData = canvas.toDataURL('image/jpeg', 0.95);
      const imgWidth = 210; // A4 width in mm
      const pageHeight = 297; // A4 height in mm
      const imgHeight = canvas.height * imgWidth / canvas.width;
      
      doc.addImage(imgData, 'JPEG', 0, 0, imgWidth, imgHeight);
      
      // Open PDF in new window
      const pdfOutput = doc.output('datauristring');
      const newWindow = window.open();
      if (!newWindow) {
        throw new Error('El navegador bloqueó la apertura de la ventana. Por favor, permita ventanas emergentes.');
      }
      
      newWindow.document.write(`
        <iframe width="100%" height="100%" src="${pdfOutput}"></iframe>
      `);
    } finally {
      document.body.removeChild(tempDiv);
    }
    
    console.log('PDF generado correctamente');
    cargando.value = false;
    setTimeout(() => {
      emit('cerrar');
    }, 2000);
  } catch (err) {
    console.error('Error al generar el PDF:', err);
    cargando.value = false;
    error.value = err.message || 'Error desconocido al generar el PDF';
  }
};

const formatNumber = (num) => {
  if (num === undefined || num === null) return '0'
  return num.toLocaleString('es-PY')
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('es-PY', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

onMounted(() => {
  console.log('Componente HTML2PDF OrdersPdf montado con orderId:', props.orderId)
  generarPDF()
})
</script>
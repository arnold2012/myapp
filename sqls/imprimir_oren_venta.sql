SELECT
    o.id_order,
    o.numero_order,
    o.fecha_expedicion,
    o.idcondicion_pago,
    d.id_order_detail,
    d.id_item,
    d.cantidad,
    cp.describe_condicion,
    i.cod_item,
    i.descripcion_item,
    i.precio_unitario,
    ii.porcentaje_iva,

    (i.precio_unitario * d.cantidad) AS total_item,

    CASE
        WHEN ii.porcentaje_iva = 10 THEN (i.precio_unitario * d.cantidad)
        ELSE 0
    END AS gravado_10,

    CASE
        WHEN ii.porcentaje_iva = 5 THEN (i.precio_unitario * d.cantidad)
        ELSE 0
    END AS gravado_5,

    CASE
        WHEN ii.porcentaje_iva = 0 THEN (i.precio_unitario * d.cantidad)
        ELSE 0
    END AS exenta,

    -- Cálculo del IVA real según porcentaje
    CASE
        WHEN ii.porcentaje_iva = 10 THEN ROUND((i.precio_unitario * d.cantidad) / 11, 2)
        WHEN ii.porcentaje_iva = 5 THEN ROUND((i.precio_unitario * d.cantidad) / 21, 2)
        ELSE 0
    END AS iva_calculado,

    -- Nueva columna para el total a pagar sin IVA
    (i.precio_unitario * d.cantidad) AS total_a_pagar

FROM
    "order" o
INNER JOIN
    order_detail d ON o.id_order = d.id_order
INNER JOIN
    condicion_pagos cp ON o.idcondicion_pago = cp.idcondicion_pago
INNER JOIN
    item i ON d.id_item = i.id_item
INNER JOIN
    item_impuesto ii ON i.id_impuesto = ii.id_impuesto;

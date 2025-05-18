procedure CREATE OR REPLACE PROCEDURE public.insert_sales_order(
    p_id_contacto bigint,
    p_numero_order numeric,
    p_fecha_expedicion date,
    p_items jsonb,
    p_idcondicion_pago bigint  -- Nuevo parámetro
)
LANGUAGE plpgsql
AS $$
DECLARE
    v_order_id bigint;
    v_item record;
BEGIN
    -- Insert into order table
    INSERT INTO public."order" (id_contacto, numero_order, fecha_expedicion, idcondicion_pago)
    VALUES (p_id_contacto, p_numero_order, p_fecha_expedicion, p_idcondicion_pago)
    RETURNING id_order INTO v_order_id;

    -- Loop through items in the JSON array and insert into order_detail
    FOR v_item IN SELECT * FROM jsonb_to_recordset(p_items) AS x(id_item bigint, cantidad integer)
    LOOP
        INSERT INTO public.order_detail (id_item, id_order, cantidad)
        VALUES (v_item.id_item, v_order_id, v_item.cantidad);
    END LOOP;

EXCEPTION
    WHEN OTHERS THEN
        RAISE EXCEPTION 'Error inserting sales order: %', SQLERRM;
END;
$$;
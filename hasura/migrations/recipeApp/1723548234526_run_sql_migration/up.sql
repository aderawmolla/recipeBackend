CREATE OR REPLACE FUNCTION public.calculate_average_rating(recipe_id uuid)
RETURNS numeric AS $$
BEGIN
    RETURN (
        SELECT AVG(r.rating)
        FROM ratings r
        WHERE r.recipe_id = recipe_id
    );
END;
$$ LANGUAGE plpgsql;

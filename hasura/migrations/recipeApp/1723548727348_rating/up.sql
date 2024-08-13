CREATE OR REPLACE FUNCTION public.calculate_average_rating(recipe_row recipes)
RETURNS numeric
LANGUAGE sql
STABLE
AS $function$
    SELECT AVG(rating) 
    FROM ratings 
    WHERE recipe_id = recipe_row.id;
$function$;

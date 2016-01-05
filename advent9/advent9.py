
def get_distance(first_city, second_city):
    if first_city in distances:
        distancemap = distances[first_city]
        if second_city in distancemap:
            return distancemap[second_city]

    if second_city in distances:
        distancemap = distances[second_city]
        if first_city in distancemap:
            return distancemap[first_city]

    raise Exception("no connection between cities found")

def calculate_tours(endpoints):
    from_endpoint = endpoints[0]
    to_endpoint = endpoints[1]
    tours = []


with open('input9.txt') as infile:
    # distances is a hashmap which maps each city to a {to_city, distance} map
    distances = {}
    cities = set()
    for line in infile:
        from_city = line.split()[0]
        to_city = line.split()[2]
        distance = line.split()[4]
        # print from_city + " " + to_city + ", " + distance
        if not from_city in distances:
            distances[from_city] = {}
        distance_map = distances[from_city]
        distance_map[to_city] = distance
        cities.add(from_city)
        cities.add(to_city)

    assert get_distance('Tambi', 'Straylight') is get_distance('Straylight', 'Tambi')

    # finding all tours:
    #    find all route endpoints
    route_endpoints = []
    fromlist = list(cities)
    for fromcity in fromlist:
        tolist = list(cities)
        tolist.remove(fromcity)
        for tocity in tolist:
            if not (fromcity,tocity) in route_endpoints and not (tocity, fromcity) in route_endpoints:
                route_endpoints.append((fromcity, tocity))

    #    pick pair of cities
    for endpoints in route_endpoints:
    #       calculate all possible tours
        tours = calculate_tours(list(endpoints))
    #       calculate each tour's costs
    # return cheapest tour

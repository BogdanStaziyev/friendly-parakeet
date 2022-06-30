# startUp
# Створити додаток для обчислення зворотньої геодезичної задачі
# _ обчислення примикання методом зєднувального трикутника 
# _ обчислення площ складу діленням на прості фігури
# _ обчислення нівелірних та теодолітних ходів з записом точок координат горизонтів і осей в БД

# Ping
http://localhost:8080/ping

# Find all: Get
			http://localhost:8080/v1/coordinates

# Find one: Get
			http://localhost:8080/v1/coordinates/{id}
			
# Create: Post
			http://localhost:8080/v1/coordinates/add
      {
    "mt": 116,
    "axis": "178 оси",
    "horizon": "1350 метр",
    "x": 25960.772,
    "y": 35375.685
}
      
      
# Update: Put(
			http://localhost:8080/v1/coordinates/update
         {
    "id": 11,
    "mt": 116,
    "axis": "178 оси",
    "horizon": "1350 метр",
    "x": 25960.772,
    "y": 35375.685
}
         
# Delete: Delete
			http://localhost:8080/v1/coordinates/{id}
      
# Invert: Get
			http://localhost:8080/v1/coordinates/{firstId}/{secondId}

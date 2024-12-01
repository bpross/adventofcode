class Brick:
    def __init__(self, start_cords, end_cords):
        self.x_start = start_cords[0]
        self.y_start = start_cords[1]
        self.z_start = start_cords[2]
        self.x_end = end_cords[0]
        self.y_end = end_cords[1]
        self.z_end = end_cords[2]

    def __str__(self):
        return "Start: " + self.x_start + "," + self.y_start + "," + self.z_start + " End: " + self.x_end + "," + self.y_end + "," + self.z_end


bricks = []

with open("sample.txt", "r") as f:
    for line in f:
        cords = line.strip().split("~")
        start_cords = cords[0].split(",")
        end_cords = cords[1].split(",")
        bricks.append(Brick(start_cords, end_cords))

for brick in bricks:
    print(brick)

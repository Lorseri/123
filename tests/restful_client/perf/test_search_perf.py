from locust import HttpUser, task, LoadTestShape
import random


class MilvusUser(HttpUser):
    host = "http://127.0.0.1:19530/v1/vector"
    vectors_to_search = [[random.random() for _ in range(768)] for _ in range(1000)]

    @task
    def search(self):
        random_id = random.randint(0, len(self.vectors_to_search) - 1)
        with self.client.post("/search",
                              json={"collectionName": "test_restful_perf",
                                    "vector": self.vectors_to_search[random_id],
                                    "filter": "id > 0"},
                              headers={"Content-Type": "application/json", "Authorization": "Bearer root:Milvus"},
                              catch_response=True
                              ) as resp:
            print(resp.status_code)
            print(resp.json()["code"])
            if resp.status_code != 200 or resp.json()["code"] != 200:
                print(resp.text)
                resp.failure("search failed")


class StagesShape(LoadTestShape):
    """
    A simple load test shape class that has different user and spawn_rate at
    different stages.

    Keyword arguments:

        stages -- A list of dicts, each representing a stage with the following keys:
            duration -- When this many seconds pass the test is advanced to the next stage
            users -- Total user count
            spawn_rate -- Number of users to start/stop per second
            stop -- A boolean that can stop that test at a specific stage

        stop_at_end -- Can be set to stop once all stages have run.
    """

    stages = [
        {"duration": 30, "users": 1, "spawn_rate": 1},
        {"duration": 60, "users": 10, "spawn_rate": 10},
        {"duration": 100, "users": 50, "spawn_rate": 10},
        {"duration": 180, "users": 100, "spawn_rate": 10},
        {"duration": 220, "users": 30, "spawn_rate": 10},
        {"duration": 230, "users": 10, "spawn_rate": 10},
        {"duration": 240, "users": 1, "spawn_rate": 1},
        {"duration": 260, "users": 1, "spawn_rate": 1},
    ]

    def tick(self):
        run_time = self.get_run_time()

        for stage in self.stages:
            if run_time < stage["duration"]:
                tick_data = (stage["users"], stage["spawn_rate"])
                return tick_data

        return None

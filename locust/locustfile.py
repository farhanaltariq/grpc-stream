from locust import HttpUser, task

class HelloWorldUser(HttpUser):
    @task
    def hello_world(self):
        self.client.get("/2024/10/22/management-system-layanan-private-booking-ladies-kesatuan-pihak-keamanan-management-bekerjasama-36")
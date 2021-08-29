from celery.result import AsyncResult
from worker_py import app

result = AsyncResult(id='52d171b5-82cd-484f-ae73-e9e1b68c4f6d', app=app)
# result.state # state
print(result.get()) # result
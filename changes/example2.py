# __base__.py

from app_config import SCHEMA_VERSION

if SCHEMA_VERSION == "v1":
    from .__base_v1__ import ValidatorBase
else:
    from .__base_v2__ import ValidatorBase

# other files
from .__base__ import ValidatorBase 

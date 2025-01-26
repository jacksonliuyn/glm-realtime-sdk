# Copyright (c) ZhiPu Corporation.
# Licensed under the MIT license.

import platform
from importlib.metadata import version


def get_user_agent():
    package_version = version("rtclient")
    python_version = platform.python_version()
    return f"zhipu-rtclient/{package_version} Python/{python_version}"

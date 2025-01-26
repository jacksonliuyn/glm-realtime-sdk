# Copyright (c) ZhiPu Corporation.
# Licensed under the MIT License.

import json
import uuid
from collections.abc import AsyncIterator
from typing import Any, Optional

from aiohttp import ClientSession, WSMsgType, WSServerHandshakeError

from rtclient.models import ServerMessageType, UserMessageType, create_message_from_dict
from rtclient.util.user_agent import get_user_agent


class ConnectionError(Exception):
    def __init__(self, message: str, headers=None):
        super().__init__(message)
        self.headers = headers


class RTLowLevelClient:
    def __init__(
        self,
        url: str,
        headers: Optional[dict[str, str]] = None,
        params: Optional[dict[str, Any]] = None,
    ):
        """初始化WebSocket客户端

        Args:
            url: WebSocket服务器地址
            headers: 请求头
            params: URL参数
        """
        self._url = url
        self._headers = headers or {}
        self._params = params or {}
        self._session = ClientSession()
        self.request_id: Optional[uuid.UUID] = None
        self.ws = None

    async def connect(self):
        """连接到WebSocket服务器"""
        try:
            self.request_id = uuid.uuid4()
            headers = {
                "User-Agent": get_user_agent(),
                **self._headers
            }
            self.ws = await self._session.ws_connect(
                self._url,
                headers=headers,
                params=self._params
            )
        except WSServerHandshakeError as e:
            await self._session.close()
            error_message = f"连接服务器失败，状态码: {e.status}"
            raise ConnectionError(error_message, e.headers) from e

    async def send(self, message: UserMessageType | dict[str, Any]):
        """发送消息到服务器

        Args:
            message: 要发送的消息，可以是 UserMessageType 或 dict
        """
        if hasattr(message, 'model_dump_json'):
            message_data = message.model_dump_json()
        else:
            message_data = json.dumps(message)
        await self.ws.send_str(message_data)

    async def send_json(self, message: dict[str, Any]):
        """发送JSON消息到服务器

        Args:
            message: 要发送的JSON消息
        """
        await self.ws.send_json(message)

    async def recv(self) -> Optional[ServerMessageType]:
        """接收服务器消息

        Returns:
            接收到的消息对象
        """
        if self.ws.closed:
            return None
        websocket_message = await self.ws.receive()
        if websocket_message.type == WSMsgType.TEXT:
            data = json.loads(websocket_message.data)
            msg = create_message_from_dict(data)
            return msg
        else:
            return None

    def __aiter__(self) -> AsyncIterator[ServerMessageType]:
        return self

    async def __anext__(self):
        message = await self.recv()
        if message is None:
            raise StopAsyncIteration
        return message

    async def close(self):
        """关闭连接"""
        if self.ws:
            await self.ws.close()
        await self._session.close()

    @property
    def closed(self) -> bool:
        """连接是否已关闭"""
        return self.ws.closed if self.ws else True

    async def __aenter__(self):
        await self.connect()
        return self

    async def __aexit__(self, *args):
        await self.close()

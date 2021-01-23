from argparse import ArgumentParser, Namespace
from copy import deepcopy
from enum import Enum, auto
from random import random
from time import sleep
from typing import Union

numeric = Union[int, float]


class Cell(Enum):
    """セルの状態"""

    ALIVE = auto()
    DEAD = auto()


class Field:
    """ライフゲームの盤面"""

    def __init__(self, height: int, width: int, init_rate: float):
        """
        コンストラクタ

        Parameters
        ----------
        height : int
            盤面の縦の長さ
        width : int
            盤面の横の長さ
        init_rate : float
            初期フィールドに生成する生存セルの割合
        """
        self.height: int = height
        self.width: int = width
        # height x widthの2次元配列の要素にセルを格納する
        # 生存セルはinit_rateの割合で格納される
        self._field: list = [
            [Cell.ALIVE if random() < init_rate else Cell.DEAD for _ in range(width)]
            for _ in range(height)
        ]

    def __str__(self) -> str:
        # 生存セルは"■"
        # 死滅セルは" "
        return "\n".join(
            "".join(["■" if cell == Cell.ALIVE else " " for cell in row])
            for row in self._field
        )

    def get_cell(self, y: int, x: int) -> Cell:
        """特定の座標のセルの状態を取得する

        Parameters
        ----------
        y : int
            Y座標
        x : int
            X座標

        Returns
        -------
        Cell
            セルの状態
        """
        return self._field[y][x]

    def set_cell(self, y: int, x: int, cell: Cell) -> None:
        """特定の座標のセルの状態を設定する

        Parameters
        ----------
        y : int
            Y座標
        x : int
            X座標
        cell : Cell
            セルの状態
        """
        self._field[y][x] = cell

    def next_step(self) -> object:
        """次の盤面を計算する

        Returns
        -------
        Field
            次の盤面
        """
        next_field = deepcopy(self)
        for y, row in enumerate(self._field):
            for x, _ in enumerate(row):
                # 周囲の生存セルの個数を計算する
                live_count = self.count_surrounding_live_cells(y, x)
                if self.get_cell(y, x) == Cell.DEAD:
                    # 現在フォーカスしているセルが死んでいる
                    if live_count == 3:
                        # 死んでいるセルに隣接する生きたセルがちょうど3つあれば、次の世代が誕生する。
                        next_field.set_cell(y, x, Cell.ALIVE)
                else:
                    # 現在フォーカスしているセルが生きている
                    if 2 <= live_count <= 3:
                        # 生きているセルに隣接する生きたセルが2つか3つならば、次の世代でも生存する。
                        next_field.set_cell(y, x, Cell.ALIVE)
                    elif live_count <= 1 or live_count >= 4:
                        # 生きているセルに隣接する生きたセルが1つ以下ならば、過疎により死滅する。
                        # 生きているセルに隣接する生きたセルが4つ以上ならば、過密により死滅する。
                        next_field.set_cell(y, x, Cell.DEAD)
        return next_field

    def count_surrounding_live_cells(self, y, x) -> int:
        """周囲の生存セルの個数を返す

        Parameters
        ----------
        y : int
            フォーカスしている座標のY座標
        x : int
            フォーカスしている座標のX座標

        Returns
        -------
        int
            生存セルの個数
        """
        count = 0
        # -1 <= _y <= 1
        for _y in range(-1, 2, 1):
            # -1 <= _x <= 1
            for _x in range(-1, 2, 1):
                if _y == 0 and _x == 0:
                    # 現在フォーカスしているセルは無視する
                    continue
                if (
                    self.get_cell(
                        (y + _y) % len(self._field), (x + _x) % len(self._field[y])
                    )
                    == Cell.ALIVE
                ):
                    # 生存セルであればカウントアップ
                    count += 1
        return count


class LifeGame():
    """ライフゲーム"""

    def __init__(self, args: Namespace):
        field_size = (args.height, args.width)
        self.curr_field: Field = Field(
            *field_size, args.init_rate
        )
        # 盤面の履歴
        self.history: list = []
        self.interval: numeric = args.interval

    def main_loop(self) -> None:
        step_count = 1
        while True:
            # 現在のステージ数を表示する
            print(f"Step {step_count}\n")
            # 盤面を見せる
            self.show()
            # 盤面が変わらなくなったら終了
            if len(self.history) > 1 and is_same_field(
                self.curr_field, self.history[0]
            ):
                print("Finish.")
                break
            # 盤面の履歴を更新
            self.update_history()
            # 次の盤面を計算する
            self.curr_field = self.curr_field.next_step()
            step_count += 1
            # 進化している間スリープ
            sleep(self.interval)
            # カーソルを左上に移動させる
            self.cursor_reset()

    def show(self) -> None:
        """コンソールに表示する"""
        print(self.curr_field)

    def update_history(self) -> None:
        """履歴を更新する"""
        if len(self.history) > 1:
            # 盤面の履歴は2つだけ持っておきたい
            self.history.pop(0)
        self.history.append(deepcopy(self.curr_field))

    def cursor_reset(self) -> None:
        """盤面の行数+α分上にカーソルを移動する"""
        print(f"\033[{self.curr_field.height + 3}A")


def is_same_field(field1: Field, field2: Field) -> bool:
    """両者の盤面が同じか判定する

    Parameters
    ----------
    field1 : Field
        ライフゲームの盤面
    field2 : Field
        ライフゲームの盤面

    Returns
    -------
    bool
        盤面が同じならTrueを返す
    """
    for y in range(field1.height):
        for x in range(field1.width):
            if field1.get_cell(y, x) != field2.get_cell(y, x):
                # 1つでも異なるセルがあれば同値ではない
                return False
    return True

def parse_command_line():
    parser = ArgumentParser()
    parser.add_argument(
        "height",
        help="Field height",
        type=int,
    )
    parser.add_argument(
        "width",
        help="Field width",
        type=int,
    )
    parser.add_argument(
        "init_rate",
        help="Percentage of surviving cells of the first generation",
        type=float,
    )
    parser.add_argument(
        "interval",
        help="Time to evolve to the next generation",
        type=float,
    )
    return parser.parse_args()

if __name__ == "__main__":
    lifegame = LifeGame(parse_command_line())
    lifegame.main_loop()

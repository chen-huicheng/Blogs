import subprocess
import case_convert 
import time

def set_clipboard(data: str):
    p = subprocess.Popen(['pbcopy'], stdin=subprocess.PIPE)
    p.stdin.write(data.encode("utf-8"))
    p.stdin.close()
    p.communicate()

def get_from_clipboard():
    p = subprocess.Popen(['pbpaste'], stdout=subprocess.PIPE)
    p.wait()
    byte_data = p.stdout.read()
    p.stdout.close()
    return byte_data.decode('utf-8')


def change_data(data: str)->str:
    return case_convert.snake_case(data)

if __name__ == "__main__":
    pre_data=""
    while True:
        # 从粘贴板拿数据
        data=get_from_clipboard()
        # 与之前不同则修改
        if data!=pre_data:
            cdata=change_data(data)
            pre_data=cdata
            set_clipboard(data=cdata)
            print(data,"-->",cdata)
        # 频控50ms
        time.sleep(0.05)


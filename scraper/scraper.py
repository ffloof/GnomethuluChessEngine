import requests
import time

def request_opening(moves):
	url = 'http://explorer.lichess.ovh/masters'

	params = dict(
		variant= "standard",
		fen= "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		play= ",".join(moves),
		since= "1952",
		until= "2022", #TODO: make sure to update when december hits
	)
	resp = requests.get(url=url, params=params)
	try:
		data = resp.json()
		print("hit", moves)
		return data["moves"]
	except:
		print("miss")
		time.sleep(5)
		return request_opening(moves)
	

def get_all_openings(current_dict, limit, move_stack=[]):
	moves = request_opening(move_stack)

	for move in moves:
		if move["white"] + move["draws"] + move["black"] < limit:
			continue
		
		next_tree = {}
		current_dict[move["uci"]] = next_tree

		move_stack.append(move["uci"])
		get_all_openings(next_tree, limit, move_stack)
		move_stack.pop()


def pretty_write_tree(file, d, layer=0, stack=[]):
	for key in d.keys():
		stack.append(key)
		file.write(" ".join(stack) + "\n")
		pretty_write_tree(file, d[key],layer+1,stack)
		stack.pop()

tree = {}
get_all_openings(tree, 1000, [])

with open("openings.txt", "w") as file:
	pretty_write_tree(file, tree)
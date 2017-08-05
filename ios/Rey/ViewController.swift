import UIKit
import Core

class ViewController: UIViewController, CoreSubscriberProtocol {

    @IBOutlet weak var textView: UITextView?
    @IBOutlet weak var button: UIButton?
    
    private func store() -> CoreStore {
        let app = UIApplication.shared.delegate as? AppDelegate;
        return app!.store
    }
    
    override func viewDidLoad() {
        button?.addTarget(self, action: #selector(self.onClick), for: .touchUpInside)
    }
    
    override func viewDidAppear(_ animated: Bool) {
        store().subscribe(self)
        CoreFetchNextPerson(store())
    }
    
    func update(_ state: CoreState!) {
        let text = getText(state: state)
        DispatchQueue.main.async {
            self.textView?.text = text
        }
    }
    
    func getText(state: CoreState) -> String {
        let name = state.currentPerson()?.name() ?? "No one"
        if (state.isFetching()) {
            return "\(name) - Loading..."
        } else {
            return name
        }
    }

    override func viewWillDisappear(_ animated: Bool) {
        store().unsubscribe(self)
    }
    
    func onClick(sender: UIButton) {
        CoreFetchNextPerson(store())
    }
    
}


import ItemList from "../ItemList";
import React, { useState, useEffect } from "react";
import agent from "../../agent";
import { connect } from "react-redux";
import { CHANGE_TAB, APPLY_SEARCH } from "../../constants/actionTypes";

const YourFeedTab = (props) => {
  if (props.token) {
    const clickHandler = (ev) => {
      ev.preventDefault();
      props.onTabClick("feed", agent.Items.feed, agent.Items.feed());
    };

    return (
      <li className="nav-item">
        <button
          type="button"
          className={props.tab === "feed" ? "nav-link active" : "nav-link"}
          onClick={clickHandler}
        >
          Your Feed
        </button>
      </li>
    );
  }
  return null;
};

const GlobalFeedTab = (props) => {
  const clickHandler = (ev) => {
    ev.preventDefault();
    props.onTabClick("all", agent.Items.all, agent.Items.all());
  };
  return (
    <li className="nav-item">
      <button
        type="button"
        className={props.tab === "all" ? "nav-link active" : "nav-link"}
        onClick={clickHandler}
      >
        Global Feed
      </button>
    </li>
  );
};

const TagFilterTab = (props) => {
  if (!props.tag) {
    return null;
  }

  return (
    <li className="nav-item">
      <button type="button" className="nav-link active">
        <i className="ion-pound"></i> {props.tag}
      </button>
    </li>
  );
};

const mapStateToProps = (state) => ({
  ...state.itemList,
  tags: state.home.tags,
  token: state.common.token,
});

const mapDispatchToProps = (dispatch) => ({
  onTabClick: (tab, pager, payload) =>
    dispatch({ type: CHANGE_TAB, tab, pager, payload }),
  onSearch: (payload) => dispatch({ type: APPLY_SEARCH, payload }),
});

const MainView = (props) => {
  const [searchTerm, setSearchTerm] = useState("");

  useEffect(() => {
    if (searchTerm.length >= 3) {
      const searchItems = async () => {
        const result = await agent.Items.search(searchTerm);
        props.onSearch(result);
      };
      searchItems();
    } else if (searchTerm === "") {
      props.onTabClick(props.tab, agent.Items.all, agent.Items.all());
    }
  }, [searchTerm]);

  const handleSearchChange = (e) => {
    setSearchTerm(e.target.value);
    props.onSearch(e.target.value);
  };

  return (
    <div>
      <div className="feed-toggle">
        <ul className="nav nav-tabs">
          <YourFeedTab
            token={props.token}
            tab={props.tab}
            onTabClick={props.onTabClick}
          />

          <GlobalFeedTab tab={props.tab} onTabClick={props.onTabClick} />

          <TagFilterTab tag={props.tag} />
        </ul>
      </div>

      <div className="search-box">
        <input
          type="text"
          id="search-box"
          placeholder="Search for items..."
          value={searchTerm}
          onChange={handleSearchChange}
        />
      </div>

      <ItemList
        pager={props.pager}
        items={props.items}
        loading={props.loading}
        itemsCount={props.itemsCount}
        currentPage={props.currentPage}
      />
    </div>
  );
};

export default connect(mapStateToProps, mapDispatchToProps)(MainView);
